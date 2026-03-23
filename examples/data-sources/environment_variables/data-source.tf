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


