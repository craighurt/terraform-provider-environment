---
layout: ""
page_title: "Provider: Environment - Functions"
description: |-
  The Environment provider maps Shell environment variables to Terraform Functions.
---

# Environment Provider - Functions

After declaring the Environment provider, you can use the `environment_variables` function to access environment variables.
The function returns a map of environment variables, optionally filtered by regex and encoded if sensitive.

## Example Usage

```terraform
terraform {
  required_providers {
    environment = {
      source  = "registry.terraform.io/craighurt/environment"
      version = "1.4.1"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.0"
    }
  }
}

provider "environment" {}

# Data source example - reads all environment variables
data "environment_variables" "all" {}

# Data source with regex filter
data "environment_variables" "regexp" {
  filter = "^LC_"
}

# Data source with filter and base64 encoding
data "environment_variables" "encoded" {
  filter    = "TOKEN"
  sensitive = true
}

# Use data source output in resources
resource "null_resource" "all" {
  triggers = data.environment_variables.all.items
}

resource "null_resource" "regexp" {
  triggers = data.environment_variables.regexp.items
}

resource "null_resource" "encoded" {
  triggers = data.environment_variables.encoded.items
}
```

## Arguments

- `filter` (optional, string): A regex pattern to filter environment variable names.
- `sensitive` (optional, bool): If true, values are base64 encoded.

## Return Type

A map of strings where keys are environment variable names and values are their values (or base64 encoded if sensitive is true).