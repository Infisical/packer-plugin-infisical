# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

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
      "echo secret_value: ${local.secrets["FOO"]}",
    ]
  }
}
