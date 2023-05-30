package crud

import (
	"context"
	"terraform-provider-crud/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CrudProvider satisfies various provider interfaces.
var _ provider.Provider = &CrudProvider{}

// CrudProvider defines the provider implementation.
type CrudProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CrudProviderModel describes the provider data model.
type CrudProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *CrudProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "crud"
	resp.Version = p.version
}

func (p *CrudProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *CrudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CrudProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	if data.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Crud API endpoint",
			"",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := data.Endpoint.ValueString()

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing CRUD API Host",
			"",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Example client configuration for data sources and resources
	client, err := client.NewClient(&host)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create CRUD API Client",
			"An unexpected error occurred when creating the Crud API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Crud Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CrudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUnicornResource,
	}
}

func (p *CrudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewunicornsDataSource,
		NewUnicornDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CrudProvider{
			version: version,
		}
	}
}
