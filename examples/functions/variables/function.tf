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
