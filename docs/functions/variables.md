---
layout: ""
page_title: "Provider: Environment - Functions"
description: |-
  Reads environment variables with optional filtering and encoding.
---

# Environment Provider - Functions

After declaring the Environment provider, you can use the `environment_variables` function to access environment variables.
The function returns a map of environment variables, optionally filtered by regex and encoded if sensitive.

## Example Usage

```terraform
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

## Arguments

- `filter` (optional, string): A regex pattern to filter environment variable names.
- `sensitive` (optional, bool): If true, values are base64 encoded.

## Return Type

A map of strings where keys are environment variable names and values are their values (or base64 encoded if sensitive is true).

### Read-Only

- `id` (String) The ID of this resource.
- `items` (Map of String)