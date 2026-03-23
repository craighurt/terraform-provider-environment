terraform {
  required_providers {
    environment = {
      source  = "registry.terraform.io/craighurt/environment"
      version = "1.4.1"
    }
  }
}

provider "environment" {}

# Ephemeral resource example - reads environment variables without storing in state
ephemeral "environment_variables" "secrets" {
  filter    = "SECRET|TOKEN"
  sensitive = true
}

# Use ephemeral values in outputs or other contexts
output "secrets" {
  value       = ephemeral.environment_variables.secrets.variables
  sensitive   = true
  description = "Sensitive environment variables (ephemeral, not stored in state)"
}

# Another ephemeral resource example with different filter
ephemeral "environment_variables" "home_vars" {
  filter    = "HOME|USER|SHELL"
  sensitive = false
}

output "user_environment" {
  value       = ephemeral.environment_variables.home_vars.variables
  description = "User-related environment variables"
}
