terraform {
  required_providers {
    environment = {
      source  = "registry.terraform.io/craighurt/environment"
      version = "1.4.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.0"
    }
  }
}

provider "environment" {}

data "environment_variables" "all" {}

data "environment_variables" "regexp" {
  filter = "^LC_"
}

data "environment_variables" "encoded" {
  filter    = "TOKEN"
  sensitive = true
}

ephemeral "environment_variables" "ephemeral_example" {
  filter = "HOME"
}

resource "null_resource" "all" {
  triggers = data.environment_variables.all.variables
}

resource "null_resource" "regexp" {
  triggers = data.environment_variables.regexp.variables
}

resource "null_resource" "encoded" {
  triggers = data.environment_variables.encoded.variables
}

