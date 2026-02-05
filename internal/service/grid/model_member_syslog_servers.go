package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberSyslogServersModel struct {
	AddressOrFqdn    types.String `tfsdk:"address_or_fqdn"`
	Certificate      types.String `tfsdk:"certificate"`
	CertificateToken types.String `tfsdk:"certificate_token"`
	ConnectionType   types.String `tfsdk:"connection_type"`
	Port             types.Int64  `tfsdk:"port"`
	LocalInterface   types.String `tfsdk:"local_interface"`
	MessageSource    types.String `tfsdk:"message_source"`
	MessageNodeId    types.String `tfsdk:"message_node_id"`
	Severity         types.String `tfsdk:"severity"`
	CategoryList     types.List   `tfsdk:"category_list"`
	OnlyCategoryList types.Bool   `tfsdk:"only_category_list"`
}

var MemberSyslogServersAttrTypes = map[string]attr.Type{
	"address_or_fqdn":    types.StringType,
	"certificate":        types.StringType,
	"certificate_token":  types.StringType,
	"connection_type":    types.StringType,
	"port":               types.Int64Type,
	"local_interface":    types.StringType,
	"message_source":     types.StringType,
	"message_node_id":    types.StringType,
	"severity":           types.StringType,
	"category_list":      types.ListType{ElemType: types.StringType},
	"only_category_list": types.BoolType,
}

var MemberSyslogServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address_or_fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The server address or FQDN.",
	},
	"certificate": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference to the underlying X509Certificate object grid:x509certificate.",
	},
	"certificate_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
	"connection_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The connection type for communicating with this server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The port this server listens on.",
	},
	"local_interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The local interface through which the appliance sends syslog messages to the syslog server.",
	},
	"message_source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source of syslog messages to be sent to the external syslog server. If set to 'INTERNAL', only messages the appliance generates will be sent to the syslog server. If set to 'EXTERNAL', the appliance sends syslog messages that it receives from other devices, such as syslog servers and routers. If set to 'ANY', the appliance sends both internal and external syslog messages.",
	},
	"message_node_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Identify the node in the syslog message.",
	},
	"severity": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The severity filter. The appliance sends log messages of the specified severity and above to the external syslog server.",
	},
	"category_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of all syslog logging categories.",
	},
	"only_category_list": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The list of selected syslog logging categories. The appliance forwards syslog messages that belong to the selected categories.",
	},
}

func ExpandMemberSyslogServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberSyslogServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberSyslogServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberSyslogServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberSyslogServers {
	if m == nil {
		return nil
	}
	to := &grid.MemberSyslogServers{
		AddressOrFqdn:    flex.ExpandStringPointer(m.AddressOrFqdn),
		CertificateToken: flex.ExpandStringPointer(m.CertificateToken),
		ConnectionType:   flex.ExpandStringPointer(m.ConnectionType),
		Port:             flex.ExpandInt64Pointer(m.Port),
		LocalInterface:   flex.ExpandStringPointer(m.LocalInterface),
		MessageSource:    flex.ExpandStringPointer(m.MessageSource),
		MessageNodeId:    flex.ExpandStringPointer(m.MessageNodeId),
		Severity:         flex.ExpandStringPointer(m.Severity),
		CategoryList:     flex.ExpandFrameworkListString(ctx, m.CategoryList, diags),
		OnlyCategoryList: flex.ExpandBoolPointer(m.OnlyCategoryList),
	}
	return to
}

func FlattenMemberSyslogServers(ctx context.Context, from *grid.MemberSyslogServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberSyslogServersAttrTypes)
	}
	m := MemberSyslogServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberSyslogServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberSyslogServersModel) Flatten(ctx context.Context, from *grid.MemberSyslogServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberSyslogServersModel{}
	}
	m.AddressOrFqdn = flex.FlattenStringPointer(from.AddressOrFqdn)
	m.Certificate = flex.FlattenStringPointer(from.Certificate)
	m.CertificateToken = flex.FlattenStringPointer(from.CertificateToken)
	m.ConnectionType = flex.FlattenStringPointer(from.ConnectionType)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.LocalInterface = flex.FlattenStringPointer(from.LocalInterface)
	m.MessageSource = flex.FlattenStringPointer(from.MessageSource)
	m.MessageNodeId = flex.FlattenStringPointer(from.MessageNodeId)
	m.Severity = flex.FlattenStringPointer(from.Severity)
	m.CategoryList = flex.FlattenFrameworkListString(ctx, from.CategoryList, diags)
	m.OnlyCategoryList = types.BoolPointerValue(from.OnlyCategoryList)
}
