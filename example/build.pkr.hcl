packer {
  required_plugins {
    infisical = {
      version = ">=0.0.1"
      source  = "github.com/infisical/infisical"
    }
  }
}

data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  project_id  = "00000000-0000-0000-0000-000000000000"

  universal_auth {
    client_id = "00000000-0000-0000-0000-000000000000"
    client_secret = "..."
  }
}

locals {
  secrets = data.infisical-secrets.dev-secrets.secrets
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo secret_key: ${local.secrets["SECRET_KEY"].secret_value}",
    ]
  }
}
