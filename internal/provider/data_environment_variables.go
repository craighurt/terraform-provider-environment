package provider

import (
	"context"
	"encoding/base64"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &EnvironmentVariablesDataSource{}

type EnvironmentVariablesDataSource struct{}

type EnvironmentVariablesDataSourceModel struct {
	Filter    types.String `tfsdk:"filter"`
	Sensitive types.Bool   `tfsdk:"sensitive"`
	Variables types.Map    `tfsdk:"variables"`
}

func NewEnvironmentVariablesDataSource() datasource.DataSource {
	return &EnvironmentVariablesDataSource{}
}

func (d *EnvironmentVariablesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_variables"
}

func (d *EnvironmentVariablesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads environment variables with optional filtering and encoding.",
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

func (d *EnvironmentVariablesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EnvironmentVariablesDataSourceModel

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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}