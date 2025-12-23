package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	gridclient "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/infoblox-nios-go-client/option"

	"github.com/infobloxopen/terraform-provider-nios/internal/service/acl"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/cloud"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/discovery"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/notification"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/service/smartfolder"
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
		dns.NewRecordAaaaResource,
		dns.NewRecordAliasResource,
		dns.NewRecordSrvResource,
		dns.NewRecordTxtResource,
		dns.NewRecordPtrResource,
		dns.NewRecordNsResource,
		dns.NewRecordDnameResource,
		dns.NewRecordCnameResource,
		dns.NewRecordMxResource,
		dns.NewRecordNaptrResource,
		dns.NewRecordTlsaResource,
		dns.NewRecordCaaResource,
		dns.NewRecordUnknownResource,
		dns.NewZoneForwardResource,
		dns.NewZoneDelegatedResource,
		dns.NewZoneAuthResource,
		dns.NewZoneRpResource,
		dns.NewViewResource,
		dns.NewZoneStubResource,
		dns.NewNsgroupResource,
		dns.NewNsgroupDelegationResource,
		dns.NewNsgroupForwardingmemberResource,
		dns.NewNsgroupForwardstubserverResource,
		dns.NewNsgroupStubmemberResource,
		dns.NewIPAllocationResource,
		dns.NewIPAssociationResource,
		dns.NewSharedrecordgroupResource,
		dns.NewSharedrecordTxtResource,
		dns.NewSharedrecordMxResource,
		dns.NewSharedrecordSrvResource,
		dns.NewSharedrecordAResource,
		dns.NewSharedrecordCnameResource,
		dns.NewSharedrecordAaaaResource,

		dhcp.NewFixedaddressResource,
		dhcp.NewSharednetworkResource,
		dhcp.NewRangeResource,
		dhcp.NewRangetemplateResource,
		dhcp.NewIpv6rangetemplateResource,
		dhcp.NewIpv6dhcpoptionspaceResource,
		dhcp.NewIpv6dhcpoptiondefinitionResource,
		dhcp.NewIpv6fixedaddresstemplateResource,
		dhcp.NewDhcpoptionspaceResource,
		dhcp.NewDhcpoptiondefinitionResource,
		dhcp.NewFixedaddresstemplateResource,
		dhcp.NewIpv6sharednetworkResource,
		dhcp.NewIpv6rangeResource,
		dhcp.NewFiltermacResource,
		dhcp.NewDhcpfailoverResource,
		dhcp.NewMacfilteraddressResource,
		dhcp.NewFilternacResource,
		dhcp.NewRoaminghostResource,

		dtc.NewDtcLbdnResource,
		dtc.NewDtcServerResource,
		dtc.NewDtcPoolResource,
		dtc.NewDtcTopologyResource,
		dtc.NewDtcMonitorSnmpResource,
		dtc.NewDtcMonitorHttpResource,
		dtc.NewDtcMonitorTcpResource,
		dtc.NewDtcMonitorPdpResource,
		dtc.NewDtcMonitorIcmpResource,
		dtc.NewDtcMonitorSipResource,
		dtc.NewDtcRecordAResource,

		ipam.NewNetworkResource,
		ipam.NewNetworktemplateResource,
		ipam.NewNetworkcontainerResource,
		ipam.NewIpv6networkcontainerResource,
		ipam.NewIpv6networkResource,
		ipam.NewNetworkviewResource,
		ipam.NewBulkhostnametemplateResource,
		ipam.NewVlanviewResource,
		ipam.NewVlanResource,
		ipam.NewVlanrangeResource,

		cloud.NewAwsrte53taskgroupResource,
		cloud.NewAwsuserResource,

		security.NewAdminuserResource,
		security.NewAdmingroupResource,
		security.NewPermissionResource,
		security.NewAdminroleResource,
		security.NewFtpuserResource,
		security.NewSnmpuserResource,
		security.NewCertificateAuthserviceResource,

		misc.NewRulesetResource,
		misc.NewBfdtemplateResource,

		smartfolder.NewSmartfolderPersonalResource,
		smartfolder.NewSmartfolderGlobalResource,

		acl.NewNamedaclResource,

		grid.NewNatgroupResource,
		grid.NewExtensibleattributedefResource,
		grid.NewUpgradegroupResource,
		grid.NewGridServicerestartGroupResource,
		grid.NewDistributionscheduleResource,

		discovery.NewDiscoveryCredentialgroupResource,
		discovery.NewVdiscoverytaskResource,

		notification.NewNotificationRuleResource,
		notification.NewNotificationRestEndpointResource,

		rpz.NewRecordRpzAResource,
		rpz.NewRecordRpzNaptrResource,
		rpz.NewRecordRpzPtrResource,
		rpz.NewRecordRpzMxResource,
		rpz.NewRecordRpzSrvResource,
		rpz.NewRecordRpzCnameResource,
		rpz.NewRecordRpzAaaaResource,
		rpz.NewRecordRpzTxtResource,
	}
}

func (p *NIOSProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{

		dns.NewRecordADataSource,
		dns.NewRecordAaaaDataSource,
		dns.NewRecordAliasDataSource,
		dns.NewRecordSrvDataSource,
		dns.NewRecordTxtDataSource,
		dns.NewRecordPtrDataSource,
		dns.NewRecordNsDataSource,
		dns.NewRecordDnameDataSource,
		dns.NewRecordCnameDataSource,
		dns.NewRecordMxDataSource,
		dns.NewRecordNaptrDataSource,
		dns.NewRecordTlsaDataSource,
		dns.NewRecordCaaDataSource,
		dns.NewRecordUnknownDataSource,
		dns.NewZoneForwardDataSource,
		dns.NewZoneDelegatedDataSource,
		dns.NewZoneAuthDataSource,
		dns.NewZoneRpDataSource,
		dns.NewViewDataSource,
		dns.NewZoneStubDataSource,
		dns.NewNsgroupDataSource,
		dns.NewNsgroupDelegationDataSource,
		dns.NewNsgroupForwardingmemberDataSource,
		dns.NewNsgroupForwardstubserverDataSource,
		dns.NewNsgroupStubmemberDataSource,
		dns.NewRecordHostDataSource,
		dns.NewSharedrecordgroupDataSource,
		dns.NewSharedrecordTxtDataSource,
		dns.NewSharedrecordMxDataSource,
		dns.NewSharedrecordSrvDataSource,
		dns.NewSharedrecordADataSource,
		dns.NewSharedrecordCnameDataSource,
		dns.NewSharedrecordAaaaDataSource,

		dhcp.NewFixedaddressDataSource,
		dhcp.NewSharednetworkDataSource,
		dhcp.NewRangetemplateDataSource,
		dhcp.NewRangeDataSource,
		dhcp.NewIpv6rangetemplateDataSource,
		dhcp.NewIpv6dhcpoptionspaceDataSource,
		dhcp.NewIpv6dhcpoptiondefinitionDataSource,
		dhcp.NewIpv6fixedaddresstemplateDataSource,
		dhcp.NewDhcpoptionspaceDataSource,
		dhcp.NewDhcpoptiondefinitionDataSource,
		dhcp.NewFixedaddresstemplateDataSource,
		dhcp.NewIpv6sharednetworkDataSource,
		dhcp.NewIpv6rangeDataSource,
		dhcp.NewFiltermacDataSource,
		dhcp.NewDhcpfailoverDataSource,
		dhcp.NewMacfilteraddressDataSource,
		dhcp.NewFilternacDataSource,
		dhcp.NewRoaminghostDataSource,

		dtc.NewDtcLbdnDataSource,
		dtc.NewDtcServerDataSource,
		dtc.NewDtcPoolDataSource,
		dtc.NewDtcTopologyRuleDataSource,
		dtc.NewDtcTopologyDataSource,
		dtc.NewDtcMonitorSnmpDataSource,
		dtc.NewDtcMonitorHttpDataSource,
		dtc.NewDtcMonitorTcpDataSource,
		dtc.NewDtcMonitorPdpDataSource,
		dtc.NewDtcMonitorIcmpDataSource,
		dtc.NewDtcMonitorSipDataSource,
		dtc.NewDtcRecordADataSource,

		ipam.NewNetworkDataSource,
		ipam.NewNetworktemplateDataSource,
		ipam.NewNetworkcontainerDataSource,
		ipam.NewIpv6networkcontainerDataSource,
		ipam.NewIpv6networkDataSource,
		ipam.NewNetworkviewDataSource,
		ipam.NewBulkhostnametemplateDataSource,
		ipam.NewVlanviewDataSource,
		ipam.NewVlanDataSource,
		ipam.NewVlanrangeDataSource,

		cloud.NewAwsrte53taskgroupDataSource,
		cloud.NewAwsuserDataSource,

		security.NewAdminroleDataSource,
		security.NewAdminuserDataSource,
		security.NewAdmingroupDataSource,
		security.NewFtpuserDataSource,
		security.NewPermissionDataSource,
		security.NewSnmpuserDataSource,
		security.NewCertificateAuthserviceDataSource,

		misc.NewRulesetDataSource,
		misc.NewBfdtemplateDataSource,

		smartfolder.NewSmartfolderPersonalDataSource,
		smartfolder.NewSmartfolderGlobalDataSource,

		acl.NewNamedaclDataSource,

		grid.NewNatgroupDataSource,
		grid.NewExtensibleattributedefDataSource,
		grid.NewUpgradegroupDataSource,
		grid.NewGridServicerestartGroupDataSource,
		grid.NewDistributionscheduleDataSource,

		discovery.NewDiscoveryCredentialgroupDataSource,
		discovery.NewVdiscoverytaskDataSource,

		notification.NewNotificationRuleDataSource,

		notification.NewNotificationRestEndpointDataSource,

		rpz.NewRecordRpzADataSource,
		rpz.NewRecordRpzNaptrDataSource,
		rpz.NewRecordRpzPtrDataSource,
		rpz.NewRecordRpzMxDataSource,
		rpz.NewRecordRpzSrvDataSource,
		rpz.NewRecordRpzCnameDataSource,
		rpz.NewRecordRpzAaaaDataSource,
		rpz.NewRecordRpzTxtDataSource,
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
	data := gridclient.Extensibleattributedef{
		Name:    gridclient.PtrString(terraformInternalIDEA),
		Type:    gridclient.PtrString("STRING"),
		Comment: gridclient.PtrString("Internal ID for Terraform Resource"),
		Flags:   gridclient.PtrString("CR"),
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
