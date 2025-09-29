package discovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/infobloxopen/infoblox-nios-go-client/discovery"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type VdiscoverytaskModel struct {
	Ref                             types.String `tfsdk:"ref"`
	AccountsList                    types.List   `tfsdk:"accounts_list"`
	AllowUnsecuredConnection        types.Bool   `tfsdk:"allow_unsecured_connection"`
	AutoConsolidateCloudEa          types.Bool   `tfsdk:"auto_consolidate_cloud_ea"`
	AutoConsolidateManagedTenant    types.Bool   `tfsdk:"auto_consolidate_managed_tenant"`
	AutoConsolidateManagedVm        types.Bool   `tfsdk:"auto_consolidate_managed_vm"`
	AutoCreateDnsHostnameTemplate   types.String `tfsdk:"auto_create_dns_hostname_template"`
	AutoCreateDnsRecord             types.Bool   `tfsdk:"auto_create_dns_record"`
	AutoCreateDnsRecordType         types.String `tfsdk:"auto_create_dns_record_type"`
	CdiscoveryFile                  types.String `tfsdk:"cdiscovery_file"`
	CdiscoveryFileToken             types.String `tfsdk:"cdiscovery_file_token"`
	Comment                         types.String `tfsdk:"comment"`
	CredentialsType                 types.String `tfsdk:"credentials_type"`
	DnsViewPrivateIp                types.String `tfsdk:"dns_view_private_ip"`
	DnsViewPublicIp                 types.String `tfsdk:"dns_view_public_ip"`
	DomainName                      types.String `tfsdk:"domain_name"`
	DriverType                      types.String `tfsdk:"driver_type"`
	EnableFilter                    types.Bool   `tfsdk:"enable_filter"`
	Enabled                         types.Bool   `tfsdk:"enabled"`
	FqdnOrIp                        types.String `tfsdk:"fqdn_or_ip"`
	GovcloudEnabled                 types.Bool   `tfsdk:"govcloud_enabled"`
	IdentityVersion                 types.String `tfsdk:"identity_version"`
	LastRun                         types.Int64  `tfsdk:"last_run"`
	Member                          types.String `tfsdk:"member"`
	MergeData                       types.Bool   `tfsdk:"merge_data"`
	MultipleAccountsSyncPolicy      types.String `tfsdk:"multiple_accounts_sync_policy"`
	Name                            types.String `tfsdk:"name"`
	NetworkFilter                   types.String `tfsdk:"network_filter"`
	NetworkList                     types.List   `tfsdk:"network_list"`
	Password                        types.String `tfsdk:"password"`
	Port                            types.Int64  `tfsdk:"port"`
	PrivateNetworkView              types.String `tfsdk:"private_network_view"`
	PrivateNetworkViewMappingPolicy types.String `tfsdk:"private_network_view_mapping_policy"`
	Protocol                        types.String `tfsdk:"protocol"`
	PublicNetworkView               types.String `tfsdk:"public_network_view"`
	PublicNetworkViewMappingPolicy  types.String `tfsdk:"public_network_view_mapping_policy"`
	RoleArn                         types.String `tfsdk:"role_arn"`
	ScheduledRun                    types.Object `tfsdk:"scheduled_run"`
	SelectedRegions                 types.String `tfsdk:"selected_regions"`
	ServiceAccountFile              types.String `tfsdk:"service_account_file"`
	ServiceAccountFileToken         types.String `tfsdk:"service_account_file_token"`
	State                           types.String `tfsdk:"state"`
	StateMsg                        types.String `tfsdk:"state_msg"`
	SyncChildAccounts               types.Bool   `tfsdk:"sync_child_accounts"`
	UpdateDnsViewPrivateIp          types.Bool   `tfsdk:"update_dns_view_private_ip"`
	UpdateDnsViewPublicIp           types.Bool   `tfsdk:"update_dns_view_public_ip"`
	UpdateMetadata                  types.Bool   `tfsdk:"update_metadata"`
	UseIdentity                     types.Bool   `tfsdk:"use_identity"`
	Username                        types.String `tfsdk:"username"`
}

var VdiscoverytaskAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"accounts_list":                       types.ListType{ElemType: types.StringType},
	"allow_unsecured_connection":          types.BoolType,
	"auto_consolidate_cloud_ea":           types.BoolType,
	"auto_consolidate_managed_tenant":     types.BoolType,
	"auto_consolidate_managed_vm":         types.BoolType,
	"auto_create_dns_hostname_template":   types.StringType,
	"auto_create_dns_record":              types.BoolType,
	"auto_create_dns_record_type":         types.StringType,
	"cdiscovery_file":                     types.StringType,
	"cdiscovery_file_token":               types.StringType,
	"comment":                             types.StringType,
	"credentials_type":                    types.StringType,
	"dns_view_private_ip":                 types.StringType,
	"dns_view_public_ip":                  types.StringType,
	"domain_name":                         types.StringType,
	"driver_type":                         types.StringType,
	"enable_filter":                       types.BoolType,
	"enabled":                             types.BoolType,
	"fqdn_or_ip":                          types.StringType,
	"govcloud_enabled":                    types.BoolType,
	"identity_version":                    types.StringType,
	"last_run":                            types.Int64Type,
	"member":                              types.StringType,
	"merge_data":                          types.BoolType,
	"multiple_accounts_sync_policy":       types.StringType,
	"name":                                types.StringType,
	"network_filter":                      types.StringType,
	"network_list":                        types.ListType{ElemType: types.StringType},
	"password":                            types.StringType,
	"port":                                types.Int64Type,
	"private_network_view":                types.StringType,
	"private_network_view_mapping_policy": types.StringType,
	"protocol":                            types.StringType,
	"public_network_view":                 types.StringType,
	"public_network_view_mapping_policy":  types.StringType,
	"role_arn":                            types.StringType,
	"scheduled_run":                       types.ObjectType{AttrTypes: VdiscoverytaskScheduledRunAttrTypes},
	"selected_regions":                    types.StringType,
	"service_account_file":                types.StringType,
	"service_account_file_token":          types.StringType,
	"state":                               types.StringType,
	"state_msg":                           types.StringType,
	"sync_child_accounts":                 types.BoolType,
	"update_dns_view_private_ip":          types.BoolType,
	"update_dns_view_public_ip":           types.BoolType,
	"update_metadata":                     types.BoolType,
	"use_identity":                        types.BoolType,
	"username":                            types.StringType,
}

var VdiscoverytaskResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"accounts_list": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The AWS Account IDs or GCP Project IDs list associated with this task.",
	},
	"allow_unsecured_connection": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Allow unsecured connection over HTTPS and bypass validation of the remote SSL certificate.",
	},
	"auto_consolidate_cloud_ea": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to insert or update cloud EAs with discovery data.",
	},
	"auto_consolidate_managed_tenant": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to replace managed tenant with discovery tenant data.",
	},
	"auto_consolidate_managed_vm": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to replace managed virtual machine with discovery vm data.",
	},
	"auto_create_dns_hostname_template": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Template string used to generate host name.",
	},
	"auto_create_dns_record": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires((path.MatchRoot("auto_create_dns_record_type"))),
			boolvalidator.AlsoRequires((path.MatchRoot("auto_create_dns_hostname_template"))),
		},
		MarkdownDescription: "Control whether to create or update DNS record using discovered data.",
	},
	"auto_create_dns_record_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("A_PTR_RECORD", "HOST_RECORD"),
		},
		MarkdownDescription: "Indicates the type of record to create if the auto create DNS record is enabled.",
	},
	"cdiscovery_file": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The path to a file containing AWS account IDs or GCP Project IDs. when multiple_accounts_sync_policy is set to UPLOAD.",
	},
	"cdiscovery_file_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The AWS account IDs or GCP Project IDs file's token.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment on the task.",
	},
	"credentials_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DIRECT", "INDIRECT"),
		},
		MarkdownDescription: "Credentials type used for connecting to the cloud management platform.",
	},
	"dns_view_private_ip": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DNS view name for private IPs.",
	},
	"dns_view_public_ip": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DNS view name for public IPs.",
	},
	"domain_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the domain to use with keystone v3.",
	},
	"driver_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("AWS", "AZURE", "GCP", "OPENSTACK", "VMWARE"),
		},
		MarkdownDescription: "Type of discovery driver.",
	},
	"enable_filter": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable filter for cloud discovery task",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Whether to enabled the cloud discovery or not.",
	},
	"fqdn_or_ip": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "FQDN or IP of the cloud management platform.",
	},
	"govcloud_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if gov cloud is enabled or disabled.",
	},
	"identity_version": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("KEYSTONE_V2", "KEYSTONE_V3"),
		},
		MarkdownDescription: "Identity service version.",
	},
	"last_run": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of last run.",
	},
	"member": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Member on which cloud discovery will be run.",
	},
	"merge_data": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to replace the old data with new or not.",
	},
	"multiple_accounts_sync_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DISCOVER", "UPLOAD"),
		},
		Default:             stringdefault.StaticString("DISCOVER"),
		MarkdownDescription: "Discover all child accounts or Upload child account ids to discover..",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of this cloud discovery task. Uniquely identify a task.",
	},
	"network_filter": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "EXCLUDE", "INCLUDE"),
		},
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "Options to filter the networks in cdiscovery task.",
	},
	"network_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "List of networks to filter in cdiscovery task.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Password used for connecting to the cloud management platform.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(443),
		MarkdownDescription: "Connection port used for connecting to the cloud management platform.",
	},
	"private_network_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Network view for private IPs.",
	},
	"private_network_view_mapping_policy": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("AUTO_CREATE", "DIRECT"),
		},
		MarkdownDescription: "Mapping policy for the network view for private IPs in discovery data.",
	},
	"protocol": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("HTTPS"),
		Validators: []validator.String{
			stringvalidator.OneOf("HTTP", "HTTPS"),
		},
		MarkdownDescription: "Connection protocol used for connecting to the cloud management platform.",
	},
	"public_network_view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Network view for public IPs.",
	},
	"public_network_view_mapping_policy": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("AUTO_CREATE", "DIRECT"),
		},
		MarkdownDescription: "Mapping policy for the network view for public IPs in discovery data.",
	},
	"role_arn": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 128),
		},
		MarkdownDescription: "Role ARN for syncing child accounts; maximum 128 characters.",
	},
	"scheduled_run": schema.SingleNestedAttribute{
		Attributes: VdiscoverytaskScheduledRunResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"selected_regions": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "String containing selected regions for discovery in comma separated format.",
	},
	"service_account_file": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The service_account_file for GCP.",
	},
	"service_account_file_token": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Service account file's token.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current state of this task.",
	},
	"state_msg": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "State message of the complete discovery process.",
	},
	"sync_child_accounts": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Synchronizing child accounts is enabled or disabled.",
	},
	"update_dns_view_private_ip": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set to true, the appliance uses a specific DNS view for private IPs.",
	},
	"update_dns_view_public_ip": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set to true, the appliance uses a specific DNS view for public IPs.",
	},
	"update_metadata": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Whether to update metadata as a result of this network discovery.",
	},
	"use_identity": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set true, all keystone connection will use \"/identity\" endpoint and port value will be ignored.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Username used for connecting to the cloud management platform.",
	},
}

func (m *VdiscoverytaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discovery.Vdiscoverytask {
	if m == nil {
		return nil
	}
	to := &discovery.Vdiscoverytask{
		AllowUnsecuredConnection:        flex.ExpandBoolPointer(m.AllowUnsecuredConnection),
		AutoConsolidateCloudEa:          flex.ExpandBoolPointer(m.AutoConsolidateCloudEa),
		AutoConsolidateManagedTenant:    flex.ExpandBoolPointer(m.AutoConsolidateManagedTenant),
		AutoConsolidateManagedVm:        flex.ExpandBoolPointer(m.AutoConsolidateManagedVm),
		AutoCreateDnsHostnameTemplate:   flex.ExpandStringPointer(m.AutoCreateDnsHostnameTemplate),
		AutoCreateDnsRecord:             flex.ExpandBoolPointer(m.AutoCreateDnsRecord),
		AutoCreateDnsRecordType:         flex.ExpandStringPointer(m.AutoCreateDnsRecordType),
		CdiscoveryFileToken:             flex.ExpandStringPointer(m.CdiscoveryFileToken),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		CredentialsType:                 flex.ExpandStringPointer(m.CredentialsType),
		DnsViewPrivateIp:                flex.ExpandStringPointer(m.DnsViewPrivateIp),
		DnsViewPublicIp:                 flex.ExpandStringPointer(m.DnsViewPublicIp),
		DomainName:                      flex.ExpandStringPointer(m.DomainName),
		DriverType:                      flex.ExpandStringPointer(m.DriverType),
		EnableFilter:                    flex.ExpandBoolPointer(m.EnableFilter),
		Enabled:                         flex.ExpandBoolPointer(m.Enabled),
		FqdnOrIp:                        flex.ExpandStringPointer(m.FqdnOrIp),
		GovcloudEnabled:                 flex.ExpandBoolPointer(m.GovcloudEnabled),
		IdentityVersion:                 flex.ExpandStringPointer(m.IdentityVersion),
		Member:                          flex.ExpandStringPointer(m.Member),
		MergeData:                       flex.ExpandBoolPointer(m.MergeData),
		MultipleAccountsSyncPolicy:      flex.ExpandStringPointer(m.MultipleAccountsSyncPolicy),
		Name:                            flex.ExpandStringPointer(m.Name),
		NetworkFilter:                   flex.ExpandStringPointer(m.NetworkFilter),
		NetworkList:                     flex.ExpandFrameworkListString(ctx, m.NetworkList, diags),
		Password:                        flex.ExpandStringPointer(m.Password),
		Port:                            flex.ExpandInt64Pointer(m.Port),
		PrivateNetworkView:              flex.ExpandStringPointer(m.PrivateNetworkView),
		PrivateNetworkViewMappingPolicy: flex.ExpandStringPointer(m.PrivateNetworkViewMappingPolicy),
		Protocol:                        flex.ExpandStringPointer(m.Protocol),
		PublicNetworkView:               flex.ExpandStringPointer(m.PublicNetworkView),
		PublicNetworkViewMappingPolicy:  flex.ExpandStringPointer(m.PublicNetworkViewMappingPolicy),
		RoleArn:                         flex.ExpandStringPointer(m.RoleArn),
		ScheduledRun:                    ExpandVdiscoverytaskScheduledRun(ctx, m.ScheduledRun, diags),
		SelectedRegions:                 flex.ExpandStringPointer(m.SelectedRegions),
		ServiceAccountFile:              flex.ExpandStringPointer(m.ServiceAccountFile),
		ServiceAccountFileToken:         flex.ExpandStringPointer(m.ServiceAccountFileToken),
		SyncChildAccounts:               flex.ExpandBoolPointer(m.SyncChildAccounts),
		UpdateDnsViewPrivateIp:          flex.ExpandBoolPointer(m.UpdateDnsViewPrivateIp),
		UpdateDnsViewPublicIp:           flex.ExpandBoolPointer(m.UpdateDnsViewPublicIp),
		UpdateMetadata:                  flex.ExpandBoolPointer(m.UpdateMetadata),
		UseIdentity:                     flex.ExpandBoolPointer(m.UseIdentity),
		Username:                        flex.ExpandStringPointer(m.Username),
	}
	return to
}

func FlattenVdiscoverytask(ctx context.Context, from *discovery.Vdiscoverytask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(VdiscoverytaskAttrTypes)
	}
	m := VdiscoverytaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, VdiscoverytaskAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *VdiscoverytaskModel) Flatten(ctx context.Context, from *discovery.Vdiscoverytask, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = VdiscoverytaskModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AccountsList = flex.FlattenFrameworkListString(ctx, from.AccountsList, diags)
	m.AllowUnsecuredConnection = types.BoolPointerValue(from.AllowUnsecuredConnection)
	m.AutoConsolidateCloudEa = types.BoolPointerValue(from.AutoConsolidateCloudEa)
	m.AutoConsolidateManagedTenant = types.BoolPointerValue(from.AutoConsolidateManagedTenant)
	m.AutoConsolidateManagedVm = types.BoolPointerValue(from.AutoConsolidateManagedVm)
	m.AutoCreateDnsHostnameTemplate = flex.FlattenStringPointer(from.AutoCreateDnsHostnameTemplate)
	m.AutoCreateDnsRecord = types.BoolPointerValue(from.AutoCreateDnsRecord)
	m.AutoCreateDnsRecordType = flex.FlattenStringPointer(from.AutoCreateDnsRecordType)
	m.CdiscoveryFileToken = flex.FlattenStringPointer(from.CdiscoveryFileToken)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CredentialsType = flex.FlattenStringPointer(from.CredentialsType)
	m.DnsViewPrivateIp = flex.FlattenStringPointer(from.DnsViewPrivateIp)
	m.DnsViewPublicIp = flex.FlattenStringPointer(from.DnsViewPublicIp)
	m.DomainName = flex.FlattenStringPointer(from.DomainName)
	m.DriverType = flex.FlattenStringPointer(from.DriverType)
	m.EnableFilter = types.BoolPointerValue(from.EnableFilter)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.FqdnOrIp = flex.FlattenStringPointer(from.FqdnOrIp)
	m.GovcloudEnabled = types.BoolPointerValue(from.GovcloudEnabled)
	m.IdentityVersion = flex.FlattenStringPointer(from.IdentityVersion)
	m.LastRun = flex.FlattenInt64Pointer(from.LastRun)
	m.Member = flex.FlattenStringPointer(from.Member)
	m.MergeData = types.BoolPointerValue(from.MergeData)
	m.MultipleAccountsSyncPolicy = flex.FlattenStringPointer(from.MultipleAccountsSyncPolicy)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkFilter = flex.FlattenStringPointer(from.NetworkFilter)
	m.NetworkList = flex.FlattenFrameworkListString(ctx, from.NetworkList, diags)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.PrivateNetworkView = flex.FlattenStringPointer(from.PrivateNetworkView)
	m.PrivateNetworkViewMappingPolicy = flex.FlattenStringPointer(from.PrivateNetworkViewMappingPolicy)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.PublicNetworkView = flex.FlattenStringPointer(from.PublicNetworkView)
	m.PublicNetworkViewMappingPolicy = flex.FlattenStringPointer(from.PublicNetworkViewMappingPolicy)
	m.RoleArn = flex.FlattenStringPointer(from.RoleArn)
	m.ScheduledRun = FlattenVdiscoverytaskScheduledRun(ctx, from.ScheduledRun, diags)
	m.SelectedRegions = flex.FlattenStringPointer(from.SelectedRegions)
	m.ServiceAccountFile = flex.FlattenStringPointer(from.ServiceAccountFile)
	m.ServiceAccountFileToken = flex.FlattenStringPointer(from.ServiceAccountFileToken)
	m.State = flex.FlattenStringPointer(from.State)
	m.StateMsg = flex.FlattenStringPointer(from.StateMsg)
	m.SyncChildAccounts = types.BoolPointerValue(from.SyncChildAccounts)
	m.UpdateDnsViewPrivateIp = types.BoolPointerValue(from.UpdateDnsViewPrivateIp)
	m.UpdateDnsViewPublicIp = types.BoolPointerValue(from.UpdateDnsViewPublicIp)
	m.UpdateMetadata = types.BoolPointerValue(from.UpdateMetadata)
	m.UseIdentity = types.BoolPointerValue(from.UseIdentity)
	m.Username = flex.FlattenStringPointer(from.Username)
}
