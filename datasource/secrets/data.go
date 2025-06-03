//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type DatasourceOutput,Config

package secrets

import (
	"fmt"
	"os"

	infisical "github.com/infisical/packer-plugin-infisical/client"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Datasource struct {
	config Config
	client *infisical.Client
}

type Config struct {
	// The host URL of your Infisical instance. If a value isn't provided, INFISICAL_HOST may be used. Default: https://app.infisical.com
	Host string `mapstructure:"host"`
	// The Infisical API Access Token. If a value isn't provided, INFISICAL_SERVICE_TOKEN may be used.
	ServiceToken string `mapstructure:"service_token"`
	// The secret path to list secrets from. Default: /
	FolderPath string `mapstructure:"folder_path"`
	// The environment to list secrets from.
	EnvSlug string `mapstructure:"env_slug" required:"true"`
}

type DatasourceOutput struct {
	Secrets map[string]string `mapstructure:"secrets"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	if d.config.FolderPath == "" {
		d.config.FolderPath = "/"
	}

	if d.config.Host == "" {
		envHost := os.Getenv("INFISICAL_HOST")
		if envHost == "" {
			d.config.Host = "https://app.infisical.com"
		} else {
			d.config.Host = envHost
		}
	}

	serviceToken := os.Getenv("INFISICAL_SERVICE_TOKEN")
	if d.config.ServiceToken == "" {
		d.config.ServiceToken = serviceToken
	}

	client, client_err := infisical.NewClient(infisical.Config{HostURL: d.config.Host, ServiceToken: d.config.ServiceToken})

	if client_err != nil {
		return client_err
	}

	d.client = client

	fmt.Println("Client")
	fmt.Println(d.client)

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	secrets, err := d.client.GetRawSecretsViaServiceToken(d.config.FolderPath, d.config.EnvSlug)
	if err != nil {
		outputSchemaType := hcldec.ImpliedType(d.OutputSpec())
		return cty.NullVal(outputSchemaType), fmt.Errorf("failed to retrieve secrets from Infisical API (folder: '%s', environment: '%s'): %w", d.config.FolderPath, d.config.EnvSlug, err)
	}

	secretsMap := make(map[string]string)
	for _, secret := range secrets {
		secretsMap[secret.SecretKey] = secret.SecretValue
	}

	output := DatasourceOutput{
		Secrets: secretsMap,
	}

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
