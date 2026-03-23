---
layout: ""
page_title: "Provider: Environment - Functions"
description: |-
  The Environment provider maps Shell environment variables to Terraform Functions.
---

# environment::variables() Function

Returns a map of environment variables from the system, optionally filtered by regex pattern and encoded as base64.

## Example Usage

```terraform
terraform {
  required_providers {
    environment = {
      source  = "registry.terraform.io/craighurt/environment"
      version = "1.4.1"
    }
  }
}

provider "environment" {}

# Get all environment variables
locals {
  all_vars = environment::variables(null, null)
}

# Get only variables matching a pattern
locals {
  aws_vars = environment::variables("^AWS_", null)
}

# Get variables with values encoded as base64
locals {
  encoded_vars = environment::variables(null, true)
}

# Get specific variables filtered and encoded
locals {
  tokens = environment::variables("_TOKEN$", true)
}

# Use in a resource
resource "local_file" "env_vars" {
  content  = jsonencode(environment::variables(null, null))
  filename = "${path.module}/env_vars.json"
}
```

## Arguments

- `filter` (required, string): A regex pattern to filter environment variable names. Pass `null` to return all variables.
- `sensitive` (required, bool): If `true`, values are base64 encoded. Pass `null` or `false` to return values as-is.

## Return Type

A map of strings where keys are environment variable names and values are their corresponding values (or base64 encoded if `sensitive` is true).