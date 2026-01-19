package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcRecordCnameModel struct {
	Ref          types.String `tfsdk:"ref"`
	AutoCreated  types.Bool   `tfsdk:"auto_created"`
	Canonical    types.String `tfsdk:"canonical"`
	Comment      types.String `tfsdk:"comment"`
	Disable      types.Bool   `tfsdk:"disable"`
	DnsCanonical types.String `tfsdk:"dns_canonical"`
	DtcServer    types.String `tfsdk:"dtc_server"`
	Ttl          types.Int64  `tfsdk:"ttl"`
	UseTtl       types.Bool   `tfsdk:"use_ttl"`
}

var DtcRecordCnameAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"auto_created":  types.BoolType,
	"canonical":     types.StringType,
	"comment":       types.StringType,
	"disable":       types.BoolType,
	"dns_canonical": types.StringType,
	"dtc_server":    types.StringType,
	"ttl":           types.Int64Type,
	"use_ttl":       types.BoolType,
}

var DtcRecordCnameResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_created": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag that indicates whether this record was automatically created by NIOS.",
	},
	"canonical": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The canonical name of the host.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"dns_canonical": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The canonical name as server by DNS protocol.",
	},
	"dtc_server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the DTC Server object with which the DTC record is associated.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time to Live (TTL) value.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ttl",
	},
}

func (m *DtcRecordCnameModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dtc.DtcRecordCname {
	if m == nil {
		return nil
	}
	to := &dtc.DtcRecordCname{
		Canonical: flex.ExpandStringPointer(m.Canonical),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Disable:   flex.ExpandBoolPointer(m.Disable),
		Ttl:       flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:    flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.DtcServer = flex.ExpandStringPointer(m.DtcServer)
	}
	return to
}

func FlattenDtcRecordCname(ctx context.Context, from *dtc.DtcRecordCname, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcRecordCnameAttrTypes)
	}
	m := DtcRecordCnameModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcRecordCnameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcRecordCnameModel) Flatten(ctx context.Context, from *dtc.DtcRecordCname, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcRecordCnameModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoCreated = types.BoolPointerValue(from.AutoCreated)
	m.Canonical = flex.FlattenStringPointer(from.Canonical)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DnsCanonical = flex.FlattenStringPointer(from.DnsCanonical)
	m.DtcServer = flex.FlattenStringPointer(from.DtcServer)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
}
