The Secrets data source displays secrets within an Infisical folder.

-> **Note:** Data sources is a feature exclusively available to HCL2 templates.

Basic examples of usage:

```hcl
data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  service_token = "st.00000000-0000-0000-0000-000000000000.d695d74bdc5c4d67ac1babd0831bd80c.b993671a3049bdd1b5f6744b44cbe0af"
}

# usage example of the data source output
locals {
  secrets = data.infisical-secrets.dev-secrets.secrets
  secret_foo_value  = secrets["FOO"].secret_value
}
```

## Configuration Reference

### Required

<!-- Code generated from the comments of the Config struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `env_slug` (string) - The environment to list secrets from.

<!-- End of code generated from the comments of the Config struct in datasource/secrets/data.go; -->

### Optional

<!-- Code generated from the comments of the Config struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `host` (string) - The host URL of your Infisical instance. If a value isn't provided, INFISICAL_HOST may be used. Default: https://app.infisical.com

- `service_token` (string) - The Infisical API Access Token. If a value isn't provided, INFISICAL_SERVICE_TOKEN may be used.

- `folder_path` (string) - The secret path to list secrets from. Default: /

<!-- End of code generated from the comments of the Config struct in datasource/secrets/data.go; -->

## Output Data

Returned secrets are in key/object pairs. Each Secret object contains data about the secret such as it's value, version, and type.

<!-- Code generated from the comments of the DatasourceOutput struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `secrets` (map[string]Secret) - Secrets

<!-- End of code generated from the comments of the DatasourceOutput struct in datasource/secrets/data.go; -->

### Secret Object

<!-- Code generated from the comments of the Secret struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `version` (int) - Version

- `workspace` (string) - Workspace

- `type` (string) - Type

- `environment` (string) - Environment

- `secret_key` (string) - Secret Key

- `secret_value` (string) - Secret Value

- `secret_comment` (string) - Secret Comment

<!-- End of code generated from the comments of the Secret struct in datasource/secrets/data.go; -->

## Authentication

Basic example of an Infisical Secrets data source authentication using service token:

```hcl
data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  service_token = "st.00000000-0000-0000-0000-000000000000.d695d74bdc5c4d67ac1babd0831bd80c.b993671a3049bdd1b5f6744b44cbe0af"
}
```
