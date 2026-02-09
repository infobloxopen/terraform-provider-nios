package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type FilterrelayagentModel struct {
	Ref                      types.String `tfsdk:"ref"`
	Uuid                     types.String `tfsdk:"uuid"`
	CircuitIdName            types.String `tfsdk:"circuit_id_name"`
	CircuitIdSubstringLength types.Int64  `tfsdk:"circuit_id_substring_length"`
	CircuitIdSubstringOffset types.Int64  `tfsdk:"circuit_id_substring_offset"`
	Comment                  types.String `tfsdk:"comment"`
	ExtAttrs                 types.Map    `tfsdk:"extattrs"`
	IsCircuitId              types.String `tfsdk:"is_circuit_id"`
	IsCircuitIdSubstring     types.Bool   `tfsdk:"is_circuit_id_substring"`
	IsRemoteId               types.String `tfsdk:"is_remote_id"`
	IsRemoteIdSubstring      types.Bool   `tfsdk:"is_remote_id_substring"`
	Name                     types.String `tfsdk:"name"`
	RemoteIdName             types.String `tfsdk:"remote_id_name"`
	RemoteIdSubstringLength  types.Int64  `tfsdk:"remote_id_substring_length"`
	RemoteIdSubstringOffset  types.Int64  `tfsdk:"remote_id_substring_offset"`
	ExtAttrsAll              types.Map    `tfsdk:"extattrs_all"`
}

var FilterrelayagentAttrTypes = map[string]attr.Type{
	"ref":                         types.StringType,
	"uuid":                        types.StringType,
	"circuit_id_name":             types.StringType,
	"circuit_id_substring_length": types.Int64Type,
	"circuit_id_substring_offset": types.Int64Type,
	"comment":                     types.StringType,
	"extattrs":                    types.MapType{ElemType: types.StringType},
	"is_circuit_id":               types.StringType,
	"is_circuit_id_substring":     types.BoolType,
	"is_remote_id":                types.StringType,
	"is_remote_id_substring":      types.BoolType,
	"name":                        types.StringType,
	"remote_id_name":              types.StringType,
	"remote_id_substring_length":  types.Int64Type,
	"remote_id_substring_offset":  types.Int64Type,
	"extattrs_all":                types.MapType{ElemType: types.StringType},
}

var FilterrelayagentResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"circuit_id_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The circuit_id_name of a DHCP relay agent filter object. This filter identifies the circuit between the remote host and the relay agent. For example, the identifier can be the ingress interface number of the circuit access unit, perhaps concatenated with the unit ID number and slot number. Also, the circuit ID can be an ATM virtual circuit ID or cable data virtual circuit ID.",
	},
	"circuit_id_substring_length": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The circuit ID substring length.",
	},
	"circuit_id_substring_offset": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The circuit ID substring offset.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a DHCP relay agent filter object.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"is_circuit_id": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MATCHES_VALUE", "ANY", "NOT_SET"),
		},
		Default:             stringdefault.StaticString("ANY"),
		MarkdownDescription: "The circuit ID matching rule of a DHCP relay agent filter object. The circuit_id value takes effect only if the value is \"MATCHES_VALUE\".",
	},
	"is_circuit_id_substring": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "Determines if the substring of circuit ID, instead of the full circuit ID, is matched.",
	},
	"is_remote_id": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MATCHES_VALUE", "ANY", "NOT_SET"),
		},
		Default:             stringdefault.StaticString("ANY"),
		MarkdownDescription: "The remote ID matching rule of a DHCP relay agent filter object. The remote_id value takes effect only if the value is Matches_Value.",
	},
	"is_remote_id_substring": schema.BoolAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "Determines if the substring of remote ID, instead of the full remote ID, is matched.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a DHCP relay agent filter object.",
	},
	"remote_id_name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The remote ID name attribute of a relay agent filter object. This filter identifies the remote host. The remote ID name can represent many different things such as the caller ID telephone number for a dial-up connection, a user name for logging in to the ISP, a modem ID, etc. When the remote ID name is defined on the relay agent, the DHCP server will have a trusted relationship to identify the remote host. The remote ID name is considered as a trusted identifier.",
	},
	"remote_id_substring_length": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The remote ID substring length.",
	},
	"remote_id_substring_offset": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The remote ID substring offset.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *FilterrelayagentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Filterrelayagent {
	if m == nil {
		return nil
	}
	to := &dhcp.Filterrelayagent{
		CircuitIdName:            flex.ExpandStringPointer(m.CircuitIdName),
		CircuitIdSubstringLength: flex.ExpandInt64Pointer(m.CircuitIdSubstringLength),
		CircuitIdSubstringOffset: flex.ExpandInt64Pointer(m.CircuitIdSubstringOffset),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		IsCircuitId:              flex.ExpandStringPointer(m.IsCircuitId),
		IsCircuitIdSubstring:     flex.ExpandBoolPointer(m.IsCircuitIdSubstring),
		IsRemoteId:               flex.ExpandStringPointer(m.IsRemoteId),
		IsRemoteIdSubstring:      flex.ExpandBoolPointer(m.IsRemoteIdSubstring),
		Name:                     flex.ExpandStringPointer(m.Name),
		RemoteIdName:             flex.ExpandStringPointer(m.RemoteIdName),
		RemoteIdSubstringLength:  flex.ExpandInt64Pointer(m.RemoteIdSubstringLength),
		RemoteIdSubstringOffset:  flex.ExpandInt64Pointer(m.RemoteIdSubstringOffset),
	}
	return to
}

func FlattenFilterrelayagent(ctx context.Context, from *dhcp.Filterrelayagent, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FilterrelayagentAttrTypes)
	}
	m := FilterrelayagentModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FilterrelayagentAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FilterrelayagentModel) Flatten(ctx context.Context, from *dhcp.Filterrelayagent, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FilterrelayagentModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.CircuitIdName = flex.FlattenStringPointer(from.CircuitIdName)
	m.CircuitIdSubstringLength = flex.FlattenInt64Pointer(from.CircuitIdSubstringLength)
	m.CircuitIdSubstringOffset = flex.FlattenInt64Pointer(from.CircuitIdSubstringOffset)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.IsCircuitId = flex.FlattenStringPointer(from.IsCircuitId)
	m.IsCircuitIdSubstring = types.BoolPointerValue(from.IsCircuitIdSubstring)
	m.IsRemoteId = flex.FlattenStringPointer(from.IsRemoteId)
	m.IsRemoteIdSubstring = types.BoolPointerValue(from.IsRemoteIdSubstring)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RemoteIdName = flex.FlattenStringPointer(from.RemoteIdName)
	m.RemoteIdSubstringLength = flex.FlattenInt64Pointer(from.RemoteIdSubstringLength)
	m.RemoteIdSubstringOffset = flex.FlattenInt64Pointer(from.RemoteIdSubstringOffset)
}
