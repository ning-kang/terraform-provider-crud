package crud

import (
	"context"
	"terraform-provider-crud/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type unicornsDataSourceModel struct {
	Unicorns []unicornsModel `tfsdk:"unicorns"`
}

// coffeesModel maps coffees schema data.
type unicornsModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Age    types.Int64  `tfsdk:"age"`
	Colour types.String `tfsdk:"colour"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &unicornsDataSource{}
	_ datasource.DataSourceWithConfigure = &unicornsDataSource{}
)

func NewunicornsDataSource() datasource.DataSource {
	return &unicornsDataSource{}
}

type unicornsDataSource struct {
	client *client.Client
}

func (d *unicornsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *unicornsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicorns"
}

func (d *unicornsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"unicorns": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"age": schema.Int64Attribute{
							Computed: true,
						},
						"colour": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *unicornsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state unicornsDataSourceModel

	unicorns, err := d.client.GetUnicorns()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Crud Unicorns",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, unicorn := range unicorns {
		unicornState := unicornsModel{
			ID:     types.StringValue(unicorn.ID),
			Name:   types.StringValue(unicorn.Name),
			Age:    types.Int64Value(int64(unicorn.Age)),
			Colour: types.StringValue(unicorn.Colour),
		}

		state.Unicorns = append(state.Unicorns, unicornState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
