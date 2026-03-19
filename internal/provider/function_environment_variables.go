package provider

import (
	"context"
	"encoding/base64"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &EnvironmentVariablesFunction{}

type EnvironmentVariablesFunction struct{}

func NewEnvironmentVariablesFunction() function.Function {
	return &EnvironmentVariablesFunction{}
}

func (f *EnvironmentVariablesFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "environment_variables"
}

func (f *EnvironmentVariablesFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Returns environment variables",
		Description: "Returns a map of environment variables, optionally filtered and encoded.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "filter",
				Description:    "Regex filter for variable names",
				AllowNullValue: true,
			},
			function.BoolParameter{
				Name:           "sensitive",
				Description:    "Whether to base64 encode values",
				AllowNullValue: true,
			},
		},
		Return: function.MapReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *EnvironmentVariablesFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var filter types.String
	var sensitive types.Bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &filter, &sensitive))
	if resp.Error != nil {
		return
	}

	filterStr := ""
	if !filter.IsNull() {
		filterStr = filter.ValueString()
	}

	sensitiveBool := false
	if !sensitive.IsNull() {
		sensitiveBool = sensitive.ValueBool()
	}

	variables := os.Environ()
	filtering := len(filterStr) > 0
	var re *regexp.Regexp
	if filtering {
		re = regexp.MustCompile(filterStr)
	}

	result := make(map[string]attr.Value)
	for _, variable := range variables {
		fields := strings.SplitN(variable, "=", 2)
		if len(fields) != 2 {
			continue
		}
		name, value := fields[0], fields[1]

		if filtering && !re.MatchString(name) {
			continue
		}
		if sensitiveBool {
			value = base64.StdEncoding.EncodeToString([]byte(value))
		}

		result[name] = types.StringValue(value)
	}

	value, _ := types.MapValue(types.StringType, result)
	resp.Result = function.NewResultData(value)
}