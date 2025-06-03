packer {
  required_plugins {
    infisical = {
      version = ">= 1.0.0"
      source  = "github.com/infisical/infisical"
    }
  }
}

data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
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
      "echo secret_key: ${local.secrets["SECRET_KEY"]}",
    ]
  }
}
