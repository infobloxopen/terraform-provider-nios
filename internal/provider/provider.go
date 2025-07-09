package provider

import (
	"context"
	"fmt"

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/service/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/service/ipam"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure NIOSProvider satisfies various provider interfaces.
var _ provider.Provider = &NIOSProvider{}

const terraformInternalIDEA = "Terraform Internal ID"

// NIOSProvider defines the provider implementation.
type NIOSProvider struct {
	version string
	commit  string
}

// NIOSProviderModel describes the provider data model.
type NIOSProviderModel struct {
	NIOSHostURL  types.String `tfsdk:"nios_host_url"`
	NIOSUsername types.String `tfsdk:"nios_username"`
	NIOSPassword types.String `tfsdk:"nios_password"`
}

func (p *NIOSProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nios"
	resp.Version = p.version
}

func (p *NIOSProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The NIOS provider is used to interact with the resources supported by Infoblox NIOS WAPI.",
		Attributes: map[string]schema.Attribute{
			"nios_host_url": schema.StringAttribute{
				Optional: true,
			},
			"nios_username": schema.StringAttribute{
				Optional: true,
			},
			"nios_password": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *NIOSProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data NIOSProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := niosclient.NewAPIClient(
		option.WithClientName(fmt.Sprintf("terraform/%s#%s", p.version, p.commit)),
		option.WithNIOSUsername(data.NIOSUsername.ValueString()),
		option.WithNIOSPassword(data.NIOSPassword.ValueString()),
		option.WithNIOSHostUrl(data.NIOSHostURL.ValueString()),
		option.WithDebug(true),
	)

	err := checkAndCreatePreRequisites(ctx, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to ensure Terraform extensible attribute exists",
			err.Error(),
		)
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *NIOSProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		dns.NewRecordAResource,
		ipam.NewNetworkResource,
	}
}

func (p *NIOSProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		dns.NewRecordADataSource,
		ipam.NewNetworkDataSource,
	}
}

func New(version, commit string) func() provider.Provider {
	return func() provider.Provider {
		return &NIOSProvider{
			version: version,
			commit:  commit,
		}
	}
}

// checkAndCreatePreRequisites creates Terraform Internal ID EA if it doesn't exist
func checkAndCreatePreRequisites(ctx context.Context, client *niosclient.APIClient) error {
	var readableAttributesForEADefinition = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"

	filters := map[string]interface{}{
		"name": terraformInternalIDEA,
	}

	apiRes, _, err := client.GridAPI.ExtensibleattributedefAPI.
		List(ctx).
		Filters(filters).
		ReturnFieldsPlus(readableAttributesForEADefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		return fmt.Errorf("error checking for existing extensible attribute: %w", err)
	}

	// If EA already exists, creation is not required
	if len(apiRes.ListExtensibleattributedefResponseObject.GetResult()) > 0 {
		return nil
	}

	// Create EA if it doesn't exist
	data := grid.Extensibleattributedef{
		Name:    grid.PtrString(terraformInternalIDEA),
		Type:    grid.PtrString("STRING"),
		Comment: grid.PtrString("Internal ID for Terraform Resource"),
		Flags:   grid.PtrString("CR"),
	}

	_, _, err = client.GridAPI.ExtensibleattributedefAPI.
		Create(ctx).
		Extensibleattributedef(data).
		ReturnFieldsPlus(readableAttributesForEADefinition).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		return fmt.Errorf("error creating Terraform extensible attribute: %w", err)
	}
	return nil
}
