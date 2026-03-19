---
layout: ""
page_title: "Ephemeral Resource: environment_variables"
description: |-
  Reads environment variables with optional filtering and encoding (ephemeral).
---

# Ephemeral Resource: environment_variables

Reads environment variables from the system with optional filtering and base64 encoding. Unlike data sources, ephemeral resources do not save data to Terraform state.

## Example Usage

```terraform
ephemeral "environment_variables" "secrets" {
  filter    = "SECRET"
  sensitive = true
}

# Use in provider configurations or other ephemeral contexts
resource "some_resource" "example" {
  config = ephemeral.environment_variables.secrets.variables
}
```

## Arguments

- `filter` (optional, string): A regex pattern to filter environment variable names.
- `sensitive` (optional, bool): If true, values are base64 encoded.

## Attributes

- `variables` (Map of String): A map of environment variables where keys are variable names and values are their values (base64 encoded if sensitive is true).