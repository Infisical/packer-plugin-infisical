# Packer Plugin Infisical

This `Infisical` multi-component plugin can be used with HashiCorp [Packer](https://developer.hashicorp.com/packer) to read secrets from Infisical.

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing automatic installation of Packer plugins. Read the [Packer documentation](https://developer.hashicorp.com/packer/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration. Then, run [`packer init`](https://developer.hashicorp.com/packer/docs/commands/init).

```hcl
packer {
  required_plugins {
    infisical = {
      source  = "github.com/infisical/infisical"
      version = ">=0.0.1"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/infisical/packer-plugin-infisical/releases).

Once you have downloaded the latest archive corresponding to your target OS, uncompress it to retrieve the plugin binary file corresponding to your platform.

To install the plugin, please follow the Packer documentation on [installing a plugin](https://developer.hashicorp.com/packer/docs/plugins#installing-plugins).

### From Sources

If you prefer to build the plugin from sources, clone the GitHub repository locally and run the command `go build` from the root directory. Upon successful compilation, a `packer-plugin-infisical` plugin binary file can be found in the root directory.

To install the compiled plugin, please follow the official Packer documentation on [installing a plugin](https://developer.hashicorp.com/packer/docs/plugins#installing-plugins).

### Configuration

For more information on how to configure the plugin, please read the documentation located in the [`docs/`](docs) directory.

## Contributing

* If you think you've found a bug in the code or you have a question regarding
  the usage of this software, please reach out to us by opening an issue in
  this GitHub repository.
* Contributions to this project are welcome: if you want to add a feature or a
  fix a bug, please do so by opening a Pull Request in this GitHub repository.
  In case of feature contribution, we kindly ask you to open an issue to
  discuss it beforehand.
