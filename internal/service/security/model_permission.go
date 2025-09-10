package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type PermissionModel struct {
	Ref          types.String `tfsdk:"ref"`
	Group        types.String `tfsdk:"group"`
	Object       types.String `tfsdk:"object"`
	Permission   types.String `tfsdk:"permission"`
	ResourceType types.String `tfsdk:"resource_type"`
	Role         types.String `tfsdk:"role"`
}

var PermissionAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"group":         types.StringType,
	"object":        types.StringType,
	"permission":    types.StringType,
	"resource_type": types.StringType,
	"role":          types.StringType,
}

var PermissionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"group": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			// stringvalidator.ExactlyOneOf(
			// 	path.MatchRoot("group"),
			// 	path.MatchRoot("role"),
			// ),
			// stringvalidator.ConflictsWith(
			// 	path.MatchRoot("object"),
			// 	path.MatchRoot("resource_type"),
			// ),
		},
		MarkdownDescription: "The name of the admin group this permission applies to.",
	},
	"object": schema.StringAttribute{
		Optional: true,
		Computed: true,
		// Validators: []validator.String{
		// 	stringvalidator.AtLeastOneOf(
		// 		path.MatchRoot("object"),
		// 		path.MatchRoot("resource_type"),
		// 	),
		// },
		MarkdownDescription: "A reference to a WAPI object, which will be the object this permission applies to.",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DENY", "READ", "WRITE"),
		},
		MarkdownDescription: "The type of permission.",
	},
	"resource_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AtLeastOneOf(
				path.MatchRoot("object"),
				path.MatchRoot("resource_type"),
			),
			stringvalidator.OneOf(
				"A", "AAAA", "AAA_EXTERNAL_SERVICE", "ADD_A_RR_WITH_EMPTY_HOSTNAME", "ALIAS", "BFD_TEMPLATE", "BULKHOST", "CAA", "CA_CERTIFICATE", "CLUSTER",
				"CNAME", "CSV_IMPORT_TASK", "DASHBOARD_TASK", "DATACOLLECTOR_CLUSTER", "DEFINED_ACL", "DELETED_OBJS_INFO_TRACKING", "DEVICE", "DHCP_FINGERPRINT", "DHCP_LEASE_HISTORY", "DHCP_MAC_FILTER",
				"DNAME", "DNS64_SYNTHESIS_GROUP", "FILE_DIST_DIRECTORY", "FIREEYE_PUBLISH_ALERT", "FIXED_ADDRESS", "FIXED_ADDRESS_TEMPLATE", "GRID_AAA_PROPERTIES", "GRID_ANALYTICS_PROPERTIES", "GRID_DHCP_PROPERTIES", "GRID_DNS_PROPERTIES",
				"GRID_FILE_DIST_PROPERTIES", "GRID_REPORTING_PROPERTIES", "GRID_SECURITY_PROPERTIES", "HOST", "HOST_ADDRESS", "HSM_GROUP", "IDNS_CERTIFICATE", "IDNS_GEO_IP", "IDNS_LBDN", "IDNS_LBDN_RECORD",
				"IDNS_MONITOR", "IDNS_POOL", "IDNS_SERVER", "IDNS_TOPOLOGY", "IMC_AVP", "IMC_PROPERTIES", "IMC_SITE", "IPV6_DHCP_LEASE_HISTORY", "IPV6_FIXED_ADDRESS", "IPV6_FIXED_ADDRESS_TEMPLATE",
				"IPV6_HOST_ADDRESS", "IPV6_NETWORK", "IPV6_NETWORK_CONTAINER", "IPV6_NETWORK_TEMPLATE", "IPV6_RANGE", "IPV6_RANGE_TEMPLATE", "IPV6_SHARED_NETWORK", "IPV6_TEMPLATE", "KERBEROS_KEY", "MEMBER",
				"MEMBER_ANALYTICS_PROPERTIES", "MEMBER_CLOUD", "MEMBER_DHCP_PROPERTIES", "MEMBER_DNS_PROPERTIES", "MEMBER_FILE_DIST_PROPERTIES", "MEMBER_SECURITY_PROPERTIES", "MSSERVER", "MS_ADSITES_DOMAIN", "MS_SUPERSCOPE", "MX",
				"NAPTR", "NETWORK", "NETWORK_CONTAINER", "NETWORK_DISCOVERY", "NETWORK_TEMPLATE", "NETWORK_VIEW", "OCSP_SERVICE", "OPTION_SPACE", "PORT_CONTROL", "PTR",
				"RANGE", "RANGE_TEMPLATE", "RECLAMATION", "REPORTING_DASHBOARD", "REPORTING_SEARCH", "RESPONSE_POLICY_RULE", "RESPONSE_POLICY_ZONE", "RESTART_SERVICE", "RESTORABLE_OPERATION", "ROAMING_HOST",
				"RULESET", "SAML_AUTH_SERVICE", "SCHEDULE_TASK", "SG_IPV4_NETWORK", "SG_IPV6_NETWORK", "SG_NETWORK_VIEW", "SHARED_A", "SHARED_AAAA", "SHARED_CNAME", "SHARED_MX",
				"SHARED_NETWORK", "SHARED_RECORD_GROUP", "SHARED_SRV", "SHARED_TXT", "SRV", "SUB_GRID", "SUB_GRID_NETWORK_VIEW_PARENT", "SUPER_HOST", "TEMPLATE", "TENANT",
				"TLSA", "TXT", "Unknown", "VIEW", "VLAN_OBJECTS", "VLAN_RANGE", "VLAN_VIEW", "ZONE",
			),
		},
		MarkdownDescription: "The type of resource this permission applies to. If 'object' is set, the permission is going to apply to child objects of the specified type, for example if 'object' was set to an authoritative zone reference and 'resource_type' was set to 'A', the permission would apply to A Resource Records within the specified zone.",
	},
	"role": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRoot("role"),
				path.MatchRoot("group"),
			),
			// stringvalidator.ConflictsWith(
			// 	path.MatchRoot("group"),
			// ),
		},
		MarkdownDescription: "The name of the role this permission applies to.",
	},
}

func ExpandPermission(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Permission {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PermissionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PermissionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Permission {
	if m == nil {
		return nil
	}
	to := &security.Permission{
		Group:        flex.ExpandStringPointer(m.Group),
		Object:       flex.ExpandStringPointer(m.Object),
		Permission:   flex.ExpandStringPointer(m.Permission),
		ResourceType: flex.ExpandStringPointer(m.ResourceType),
		Role:         flex.ExpandStringPointer(m.Role),
	}
	return to
}

func FlattenPermission(ctx context.Context, from *security.Permission, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PermissionAttrTypes)
	}
	m := PermissionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PermissionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PermissionModel) Flatten(ctx context.Context, from *security.Permission, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PermissionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Group = flex.FlattenStringPointer(from.Group)
	m.Object = flex.FlattenStringPointer(from.Object)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.ResourceType = flex.FlattenStringPointer(from.ResourceType)
	m.Role = flex.FlattenStringPointer(from.Role)
}
