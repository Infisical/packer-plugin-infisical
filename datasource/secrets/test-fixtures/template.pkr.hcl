# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "infisical-secrets" "dev-secrets" {
  folder_path = "/"
  env_slug    = "dev"
  host        = "http://localhost:8080"
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
      "echo secret_value: ${local.secrets["FOO"].secret_value}",
    ]
  }
}
