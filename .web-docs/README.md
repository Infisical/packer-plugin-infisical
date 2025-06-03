### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    infisical = {
      # source represents the GitHub URI to the plugin repository without the `packer-plugin-` prefix.
      source  = "github.com/infisical/infisical"
      version = ">=0.0.1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/infisical/infisical
```

### Components

#### Data Sources

- [infisical-secrets](/packer/integrations/infisical/infisical/latest/components/datasource/secrets) - Retrieve secrets from a folder.

### Authentication

The Infisical provider currently supports these authentication methods:

- Service Token

#### Service Token

The service token can either be provided via `INFISICAL_SERVICE_TOKEN` environment variable or manually passed as shown below:

```hcl
data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  service_token = "st.00000000-0000-0000-0000-000000000000.d695d74bdc5c4d67ac1babd0831bd80c.b993671a3049bdd1b5f6744b44cbe0af"
}
```
