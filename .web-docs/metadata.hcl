# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Infisical"
  description = "The Infisical plugin can be used with HashiCorp Packer to read secrets from Infisical."
  identifier = "packer/infisical/infisical"
  # flags = ["hcp-ready"]
  component {
    type = "data-source"
    name = "Secrets"
    slug = "secrets"
  }
}
