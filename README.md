# Terraform Provider Environment [![release](https://github.com/craighurt/terraform-provider-environment/actions/workflows/release.yml/badge.svg)](https://github.com/craighurt/terraform-provider-environment/actions/workflows/release.yml)

Terraform provider able to detect environment settings.
Useful for debugging terraform running in CI.

## Test

```shell
make test
make testacc
```

## Build

Run the following command to build the provider

```shell
make build
```

## Install

```shell
make install
```

## Example

### Using Provider Function

```hcl
provider "environment" {}

locals {
  all     = provider::environment::environment_variables(null, null)
  regexp  = provider::environment::environment_variables("^LC_", null)
  encoded = provider::environment::environment_variables("TOKEN", true)
}

resource "null_resource" "all" {
  triggers = local.all
}

resource "null_resource" "regexp" {
  triggers = local.regexp
}

resource "null_resource" "encoded" {
  triggers = local.encoded
}
```

### Using Ephemeral Resource

```hcl
provider "environment" {}

ephemeral "environment_variables" "secrets" {
  filter    = "SECRET"
  sensitive = true
}

# Use ephemeral values in provider configurations or other ephemeral contexts
resource "some_resource" "example" {
  # Ephemeral values can be used here if the attribute supports ephemeral
  config = ephemeral.environment_variables.secrets.variables
}
```

The example code is available inside example directory.

```shell
terraform init && terraform plan
```

```shell
Terraform will perform the following actions:

  # null_resource.all will be created
  + resource "null_resource" "all" {
      + id       = (known after apply)
      + triggers = {
          + "PWD"                                 = "/terraform/terraform-provider-environment/examples"
          + "TERM"                                = "xterm-256color"
          + "SHELL"                               = "/bin/zsh"
          + "SHLVL"                               = "1"
          [...]
    }

  # null_resource.encoded will be created
  + resource "null_resource" "encoded" {
      + id       = (known after apply)
      + triggers = {
          + "TFE_TOKEN" = "ZXhhbXBsZS5hdGxhc3YxLnNlY3JldHRva2Vu"
        }
    }

  # null_resource.regexp will be created
  + resource "null_resource" "regexp" {
      + id       = (known after apply)
      + triggers = {
          + "LC_CTYPE"            = "UTF-8"
          + "LC_TERMINAL"         = "iTerm2"
          + "LC_TERMINAL_VERSION" = "3.3.11"
        }
    }

Plan: 3 to add, 0 to change, 0 to destroy.

```
