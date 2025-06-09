The Secrets data source displays secrets within an Infisical folder.

-> **Note:** Data sources is a feature exclusively available to HCL2 templates.

Basic examples of usage:

```hcl
data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  project_id  = "00000000-0000-0000-0000-000000000000"

  universal_auth {
    client_id = "00000000-0000-0000-0000-000000000000"
    client_secret = "..."
  }
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

- `project_id` (string) - The project to list secrets from.

- `env_slug` (string) - The environment to list secrets from.

<!-- End of code generated from the comments of the Config struct in datasource/secrets/data.go; -->

### Optional

<!-- Code generated from the comments of the Config struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `host` (string) - The host URL of your Infisical instance. If a value isn't provided, INFISICAL_HOST may be used. Default: https://app.infisical.com

- `folder_path` (string) - The secret path to list secrets from. Default: /

- `universal_auth` (UniversalAuth) - Configuration for Infisical Universal Authentication.

<!-- End of code generated from the comments of the Config struct in datasource/secrets/data.go; -->

### Universal Auth Object

**Required:**

<!-- Code generated from the comments of the UniversalAuth struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `client_id` (string) - The Client ID for Infisical Universal Authentication.

<!-- End of code generated from the comments of the UniversalAuth struct in datasource/secrets/data.go; -->

**Optional:**

<!-- Code generated from the comments of the UniversalAuth struct in datasource/secrets/data.go; DO NOT EDIT MANUALLY -->

- `client_secret` (string) - The Client Secret for Infisical Universal Authentication. You may use INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET env variable instead.

<!-- End of code generated from the comments of the UniversalAuth struct in datasource/secrets/data.go; -->

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

Basic example of an Infisical Secrets data source authentication using universal auth:

```hcl
data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  project_id  = "00000000-0000-0000-0000-000000000000"

  universal_auth {
    client_id = "00000000-0000-0000-0000-000000000000"
    client_secret = "..."
  }
}
```

`client_secret` may be left blank if you're using the `INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET` environment variable.
