package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccEnvironmentVariablesFunction(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"environment": providerserver.NewProtocol6WithError(&EnvironmentProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "environment" {}

				locals {
					env_vars = provider::environment::environment_variables(null, null)
				}

				output "env_vars" {
					value = local.env_vars
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						// Since we can't predict exact env vars, just check it's not empty
						output := s.RootModule().Outputs["env_vars"]
						if output.Value == nil || output.Value == "" {
							return fmt.Errorf("expected non-empty output")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccEnvironmentVariablesFunctionWithFilter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"environment": providerserver.NewProtocol6WithError(&EnvironmentProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "environment" {}

				locals {
					env_vars = provider::environment::environment_variables("HOME", null)
				}

				output "env_vars" {
					value = local.env_vars
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						output := s.RootModule().Outputs["env_vars"]
						if output.Value == nil {
							return fmt.Errorf("expected output")
						}
						// Check that HOME is present
						outputStr := fmt.Sprintf("%v", output.Value)
						if !strings.Contains(outputStr, "HOME") {
							return fmt.Errorf("expected HOME in output")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccEnvironmentVariablesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"environment": providerserver.NewProtocol6WithError(&EnvironmentProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "environment" {}

				data "environment_variables" "test" {
					filter = "HOME"
				}

				output "env_vars" {
					value = data.environment_variables.test.variables
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						output := s.RootModule().Outputs["env_vars"]
						if output.Value == nil {
							return fmt.Errorf("expected output")
						}
						// Check that HOME is present
						outputStr := fmt.Sprintf("%v", output.Value)
						if !strings.Contains(outputStr, "HOME") {
							return fmt.Errorf("expected HOME in output")
						}
						return nil
					},
				),
			},
		},
	})
}