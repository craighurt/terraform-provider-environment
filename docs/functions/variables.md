---
page_title: "variables Function - terraform-provider-environment"
subcategory: ""
description: |-
  Returns environment variables, optionally filtered and encoded.
---

# variables Function

Returns a map of environment variables from the system, optionally filtered by regex and encoded as base64.

## Example Usage

```terraform
terraform {
  required_providers {
    environment = {
      source  = "registry.terraform.io/craighurt/environment"
      version = "1.4.2"
    }
  }
}

provider "environment" {}

# Get all environment variables
locals {
  all_vars = provider::environment::variables(null, null)
}

# Get only variables matching a pattern
locals {
  aws_vars = provider::environment::variables("^AWS_", null)
}

# Get variables with values encoded as base64
locals {
  encoded_vars = provider::environment::variables(null, true)
}

# Get specific variables filtered and encoded
locals {
  tokens = provider::environment::variables("_TOKEN$", true)
}

# Use function results in a resource
output "all_environment_variables" {
  value       = provider::environment::variables(null, null)
  description = "All environment variables available to Terraform"
}

output "aws_environment_variables" {
  value       = provider::environment::variables("^AWS_", null)
  description = "Only AWS-related environment variables"
}

output "base64_encoded_tokens" {
  value       = provider::environment::variables("_TOKEN$", true)
  sensitive   = true
  description = "Base64-encoded token environment variables"
}
```

## Arguments

- `filter` (required, string): A regex pattern to filter environment variable names. Pass `null` to return all variables.
- `sensitive` (required, bool): If `true`, values are base64 encoded. Pass `null` or `false` to return values as-is.

## Return Type

A map of strings where keys are environment variable names and values are their values, or base64-encoded values when `sensitive` is `true`.
