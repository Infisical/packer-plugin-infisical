//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type DatasourceOutput,Config,Secret,UniversalAuth

package secrets

import (
	"fmt"
	"os"

	infisical "github.com/infisical/packer-plugin-infisical/client"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Datasource struct {
	config Config
	client *infisical.Client
}

type UniversalAuth struct {
	// The Client ID for Infisical Universal Authentication.
	ClientID string `mapstructure:"client_id" required:"true"`
	// The Client Secret for Infisical Universal Authentication. You may use INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET env variable instead.
	ClientSecret string `mapstructure:"client_secret"`
}

type Config struct {
	// The host URL of your Infisical instance. If a value isn't provided, INFISICAL_HOST may be used. Default: https://app.infisical.com
	Host string `mapstructure:"host"`
	// The project to list secrets from.
	ProjectId string `mapstructure:"project_id" required:"true"`
	// The secret path to list secrets from. Default: /
	FolderPath string `mapstructure:"folder_path"`
	// The environment to list secrets from.
	EnvSlug string `mapstructure:"env_slug" required:"true"`

	// Configuration for Infisical Universal Authentication.
	UniversalAuth UniversalAuth `mapstructure:"universal_auth"`
}

type Secret struct {
	Version       int    `mapstructure:"version" required:"true"`
	Workspace     string `mapstructure:"workspace" required:"true"`
	Type          string `mapstructure:"type" required:"true"`
	Environment   string `mapstructure:"environment" required:"true"`
	SecretKey     string `mapstructure:"secret_key" required:"true"`
	SecretValue   string `mapstructure:"secret_value" required:"true"`
	SecretComment string `mapstructure:"secret_comment" required:"true"`
}

type DatasourceOutput struct {
	Secrets map[string]Secret `mapstructure:"secrets" required:"true"`
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

	// Validate UniversalAuth credentials
	if d.config.UniversalAuth.ClientID == "" {
		return fmt.Errorf("universal_auth.client_id is required")
	}
	if d.config.UniversalAuth.ClientSecret == "" {
		clientSecret := os.Getenv("INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET")
		if clientSecret == "" {
			return fmt.Errorf("universal_auth.client_secret config variable or INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET env variable required")
		}
		d.config.UniversalAuth.ClientSecret = clientSecret
	}

	clientCfg := infisical.Config{
		HostURL:      d.config.Host,
		ClientId:     d.config.UniversalAuth.ClientID,
		ClientSecret: d.config.UniversalAuth.ClientSecret,
	}

	client, client_err := infisical.NewClient(clientCfg)
	if client_err != nil {
		return fmt.Errorf("failed to initialize Infisical client using Universal Authentication: %w", client_err)
	}

	d.client = client

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	secrets, err := d.client.GetRawSecrets(d.config.FolderPath, d.config.EnvSlug, d.config.ProjectId)
	if err != nil {
		outputSchemaType := hcldec.ImpliedType(d.OutputSpec())
		return cty.NullVal(outputSchemaType), fmt.Errorf("failed to retrieve secrets from Infisical API (folder: '%s', environment: '%s'): %w", d.config.FolderPath, d.config.EnvSlug, err)
	}

	ctySecretsMap := make(map[string]cty.Value)
	for _, s := range secrets {
		secretVal := cty.ObjectVal(map[string]cty.Value{
			"version":        cty.NumberIntVal(int64(s.Version)),
			"workspace":      cty.StringVal(s.Workspace),
			"type":           cty.StringVal(s.Type),
			"environment":    cty.StringVal(s.Environment),
			"secret_key":     cty.StringVal(s.SecretKey),
			"secret_value":   cty.StringVal(s.SecretValue),
			"secret_comment": cty.StringVal(s.SecretComment),
		})
		ctySecretsMap[s.SecretKey] = secretVal
	}

	outputVal := cty.ObjectVal(map[string]cty.Value{
		"secrets": cty.MapVal(ctySecretsMap),
	})

	return outputVal, nil
}
