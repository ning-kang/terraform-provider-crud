package crud

import (
	"context"
	"encoding/json"
	"terraform-provider-crud/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type unicornResourceModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Age    types.Int64  `tfsdk:"age"`
	Colour types.String `tfsdk:"colour"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &unicornResource{}
	_ resource.ResourceWithConfigure   = &unicornResource{}
	_ resource.ResourceWithImportState = &unicornResource{}
)

// NewunicornResource is a helper function to simplify the provider implementation.
func NewUnicornResource() resource.Resource {
	return &unicornResource{}
}

// unicornResource is the resource implementation.
type unicornResource struct {
	client *client.Client
}

// Configure adds the provider configured client to the resource.
func (r *unicornResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

// Metadata returns the resource type name.
func (r *unicornResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicorn"
}

// Schema defines the schema for the resource.
func (r *unicornResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"age": schema.Int64Attribute{
				Required: true,
			},
			"colour": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *unicornResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan unicornResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	item := client.UnicornItem{
		Name:   plan.Name.ValueString(),
		Age:    int(plan.Age.ValueInt64()),
		Colour: plan.Colour.ValueString(),
	}

	// Create new unicorn
	unicorn, err := r.client.CreateUnicorn(&item)
	//str, _ := json.Marshal(&item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating unicorn",
			"Could not create unicorn, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan = unicornResourceModel{
		ID:     types.StringValue(unicorn.ID),
		Name:   types.StringValue(unicorn.Name),
		Age:    types.Int64Value(int64(unicorn.Age)),
		Colour: types.StringValue(unicorn.Colour),
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *unicornResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state unicornResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	unicorn, err := r.client.GetUnicorn(state.ID.ValueString())
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
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *unicornResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan unicornResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	item := client.UnicornItem{
		Name:   plan.Name.ValueString(),
		Age:    int(plan.Age.ValueInt64()),
		Colour: plan.Colour.ValueString(),
	}

	// Update existing unicorn
	_, err := r.client.UpdateUnicorn(plan.ID.ValueString(), &item)
	str, _ := json.Marshal(&item)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Crud unicorn",
			"Could not update unicorn, unexpected error: "+err.Error()+string(str),
		)
		return
	}

	// Fetch updated items from GetUnicorn as UpdateUnicorn items are not
	// populated.
	unicorn, err := r.client.GetUnicorn(plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Crud unicorn",
			"Could not read Crud unicorn ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Update resource state with updated items and timestamp
	plan = unicornResourceModel{
		ID:     types.StringValue(plan.ID.ValueString()),
		Name:   types.StringValue(unicorn.Name),
		Age:    types.Int64Value(int64(unicorn.Age)),
		Colour: types.StringValue(unicorn.Colour),
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *unicornResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state unicornResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	err := r.client.DeleteUnicorn(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Crud unicorn",
			"Could not delete unicorn, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *unicornResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
