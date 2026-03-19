package provider

import (
	"context"
	"encoding/base64"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ephemeral.EphemeralResource = &EnvironmentVariablesEphemeralResource{}

type EnvironmentVariablesEphemeralResource struct{}

type EnvironmentVariablesEphemeralModel struct {
	Filter    types.String `tfsdk:"filter"`
	Sensitive types.Bool   `tfsdk:"sensitive"`
	Variables types.Map    `tfsdk:"variables"`
}

func NewEnvironmentVariablesEphemeralResource() ephemeral.EphemeralResource {
	return &EnvironmentVariablesEphemeralResource{}
}

func (e *EnvironmentVariablesEphemeralResource) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_variables"
}

func (e *EnvironmentVariablesEphemeralResource) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads environment variables with optional filtering and encoding (ephemeral).",
		Attributes: map[string]schema.Attribute{
			"filter": schema.StringAttribute{
				Description: "Regex filter for variable names",
				Optional:    true,
			},
			"sensitive": schema.BoolAttribute{
				Description: "Whether to base64 encode values",
				Optional:    true,
			},
			"variables": schema.MapAttribute{
				Description: "Map of environment variables",
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}

func (e *EnvironmentVariablesEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data EnvironmentVariablesEphemeralModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filterStr := ""
	if !data.Filter.IsNull() {
		filterStr = data.Filter.ValueString()
	}

	sensitiveBool := false
	if !data.Sensitive.IsNull() {
		sensitiveBool = data.Sensitive.ValueBool()
	}

	variables := os.Environ()
	filtering := len(filterStr) > 0
	var re *regexp.Regexp
	if filtering {
		var err error
		re, err = regexp.Compile(filterStr)
		if err != nil {
			resp.Diagnostics.AddError("Invalid filter regex", err.Error())
			return
		}
	}

	result := make(map[string]string)
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

		result[name] = value
	}

	mapValue, diags := types.MapValueFrom(ctx, types.StringType, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Variables = mapValue
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}