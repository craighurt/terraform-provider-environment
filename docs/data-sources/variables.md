---
layout: ""
page_title: "Data Source: environment_variables"
description: |-
  Reads environment variables with optional filtering and encoding.
---

# Data Source: environment_variables

Reads environment variables from the system, with optional filtering by regex and base64 encoding for sensitive values.

## Example Usage

```terraform
data "environment_variables" "all" {}

data "environment_variables" "filtered" {
  filter = "^MY_"
}

data "environment_variables" "encoded" {
  filter    = "SECRET"
  sensitive = true
}

output "all_vars" {
  value = data.environment_variables.all.variables
}
```

## Arguments

- `filter` (optional, string): A regex pattern to filter environment variable names.
- `sensitive` (optional, bool): If true, values are base64 encoded.

## Attributes

- `variables` (Map of String): A map of environment variables where keys are variable names and values are their values (base64 encoded if sensitive is true).