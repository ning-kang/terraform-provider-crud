package crud

import (
	"context"
	"terraform-provider-crud/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &unicornDataSource{}
	_ datasource.DataSourceWithConfigure = &unicornDataSource{}
)

// NewCoffeesDataSource is a helper function to simplify the provider implementation.
func NewUnicornDataSource() datasource.DataSource {
	return &unicornDataSource{}
}

// coffeesDataSource is the data source implementation.
type unicornDataSource struct {
	client *client.Client
}

// Configure adds the provider configured client to the data source.
func (d *unicornDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

// Metadata returns the data source type name.
func (d *unicornDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicorn"
}

// Schema defines the schema for the data source.
func (d *unicornDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *unicornDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state unicornResourceModel

	// Get refreshed order value from HashiCups
	unicorn, err := d.client.GetUnicorn(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Crud Unicorn",
			"Could not read Crud Unicorn ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state = unicornResourceModel{
		ID:     types.StringValue(unicorn.ID),
		Name:   types.StringValue(unicorn.Name),
		Age:    types.Int64Value(int64(unicorn.Age)),
		Colour: types.StringValue(unicorn.Colour),
	}

	// Set refreshed state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
