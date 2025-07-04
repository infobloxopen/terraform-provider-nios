package dtc

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type DtcServerModel struct {
	Ref                  types.String `tfsdk:"ref"`
	AutoCreateHostRecord types.Bool   `tfsdk:"auto_create_host_record"`
	Comment              types.String `tfsdk:"comment"`
	Disable              types.Bool   `tfsdk:"disable"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	Health               types.Object `tfsdk:"health"`
	Host                 types.String `tfsdk:"host"`
	Monitors             types.List   `tfsdk:"monitors"`
	Name                 types.String `tfsdk:"name"`
	SniHostname          types.String `tfsdk:"sni_hostname"`
	UseSniHostname       types.Bool   `tfsdk:"use_sni_hostname"`
}

var DtcServerAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"auto_create_host_record": types.BoolType,
	"comment":                 types.StringType,
	"disable":                 types.BoolType,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"extattrs_all":            types.MapType{ElemType: types.StringType},
	"health":                  types.ObjectType{AttrTypes: DtcServerHealthAttrTypes},
	"host":                    types.StringType,
	"monitors":                types.ListType{ElemType: types.ObjectType{AttrTypes: DtcServerMonitorsAttrTypes}},
	"name":                    types.StringType,
	"sni_hostname":            types.StringType,
	"use_sni_hostname":        types.BoolType,
}

var DtcServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_create_host_record": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the DTC Server; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"health": schema.SingleNestedAttribute{
		Attributes: DtcServerHealthResourceSchemaAttributes,
		Computed:   true,
	},
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The address or FQDN of the server.",
	},
	"monitors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DtcServerMonitorsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of IP/FQDN and monitor pairs to be used for additional monitoring.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The DTC Server display name.",
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Name should not have leading or trailing whitespace",
			),
		},
	},
	"sni_hostname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_sni_hostname")),
		},
		MarkdownDescription: "The hostname for Server Name Indication (SNI) in FQDN format.",
	},
	"use_sni_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: sni_hostname",
	},
}

func ExpandDtcServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcServer {
	if m == nil {
		return nil
	}
	to := &dtc.DtcServer{
		Ref:                  flex.ExpandStringPointer(m.Ref),
		AutoCreateHostRecord: flex.ExpandBoolPointer(m.AutoCreateHostRecord),
		Comment:              flex.ExpandStringPointer(m.Comment),
		Disable:              flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:             ExpandExtAttr(ctx, m.ExtAttrs, diags),
		Health:               ExpandDtcServerHealth(ctx, m.Health, diags),
		Host:                 flex.ExpandStringPointer(m.Host),
		Monitors:             flex.ExpandFrameworkListNestedBlock(ctx, m.Monitors, diags, ExpandDtcServerMonitors),
		Name:                 flex.ExpandStringPointer(m.Name),
		SniHostname:          flex.ExpandStringPointer(m.SniHostname),
		UseSniHostname:       flex.ExpandBoolPointer(m.UseSniHostname),
	}
	return to
}

func FlattenDtcServer(ctx context.Context, from *dtc.DtcServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcServerAttrTypes)
	}
	m := DtcServerModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, DtcServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcServerModel) Flatten(ctx context.Context, from *dtc.DtcServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcServerModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoCreateHostRecord = types.BoolPointerValue(from.AutoCreateHostRecord)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrsAll = FlattenExtAttr(ctx, from.ExtAttrs, diags)
	m.Health = FlattenDtcServerHealth(ctx, from.Health, diags)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Monitors = flex.FlattenFrameworkListNestedBlock(ctx, from.Monitors, DtcServerMonitorsAttrTypes, diags, FlattenDtcServerMonitors)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SniHostname = flex.FlattenStringPointer(from.SniHostname)
	m.UseSniHostname = types.BoolPointerValue(from.UseSniHostname)
}
