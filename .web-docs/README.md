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

- Universal Auth

#### Universal Auth

Usage example:

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
