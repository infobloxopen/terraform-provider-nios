package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcMonitorSipModel struct {
	Ref          types.String `tfsdk:"ref"`
	Uuid         types.String `tfsdk:"uuid"`
	Ciphers      types.String `tfsdk:"ciphers"`
	ClientCert   types.String `tfsdk:"client_cert"`
	Comment      types.String `tfsdk:"comment"`
	ExtAttrs     types.Map    `tfsdk:"extattrs"`
	Interval     types.Int64  `tfsdk:"interval"`
	Name         types.String `tfsdk:"name"`
	Port         types.Int64  `tfsdk:"port"`
	Request      types.String `tfsdk:"request"`
	Result       types.String `tfsdk:"result"`
	ResultCode   types.Int64  `tfsdk:"result_code"`
	RetryDown    types.Int64  `tfsdk:"retry_down"`
	RetryUp      types.Int64  `tfsdk:"retry_up"`
	Timeout      types.Int64  `tfsdk:"timeout"`
	Transport    types.String `tfsdk:"transport"`
	ValidateCert types.Bool   `tfsdk:"validate_cert"`
	ExtAttrsAll  types.Map    `tfsdk:"extattrs_all"`
}

var DtcMonitorSipAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"uuid":          types.StringType,
	"ciphers":       types.StringType,
	"client_cert":   types.StringType,
	"comment":       types.StringType,
	"extattrs":      types.MapType{ElemType: types.StringType},
	"interval":      types.Int64Type,
	"name":          types.StringType,
	"port":          types.Int64Type,
	"request":       types.StringType,
	"result":        types.StringType,
	"result_code":   types.Int64Type,
	"retry_down":    types.Int64Type,
	"retry_up":      types.Int64Type,
	"timeout":       types.Int64Type,
	"transport":     types.StringType,
	"validate_cert": types.BoolType,
	"extattrs_all":  types.MapType{ElemType: types.StringType},
}

var DtcMonitorSipResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"ciphers": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "An optional cipher list for secure TLS/SIPS connection.",
	},
	"client_cert": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "An optional client certificate, supplied in TLS and SIPS mode if present.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
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
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for TCP health check.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name for this DTC monitor.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5060),
		MarkdownDescription: "The port value for SIP requests.",
	},
	"request": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A SIP request to send",
	},
	"result": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("CODE_IS"),
		Validators: []validator.String{
			stringvalidator.OneOf("CODE_IS", "CODE_IS_NOT", "ANY"),
		},
		MarkdownDescription: "The type of an expected result.",
	},
	"result_code": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(200),
		MarkdownDescription: "The expected return code value.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as down to be treated as dead after it was alive.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as up to be treated as alive after it was dead.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for TCP health check in seconds.",
	},
	"transport": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("TCP"),
		Validators: []validator.String{
			stringvalidator.OneOf("UDP", "TCP", "TLS", "SIPS"),
		},
		MarkdownDescription: "The transport layer protocol to use for SIP check.",
	},
	"validate_cert": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the validation of the remote server's certificate is enabled.",
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

func (m *DtcMonitorSipModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcMonitorSip {
	if m == nil {
		return nil
	}
	to := &dtc.DtcMonitorSip{
		Ciphers:      flex.ExpandStringPointer(m.Ciphers),
		ClientCert:   flex.ExpandStringPointer(m.ClientCert),
		Comment:      flex.ExpandStringPointer(m.Comment),
		ExtAttrs:     ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Interval:     flex.ExpandInt64Pointer(m.Interval),
		Name:         flex.ExpandStringPointer(m.Name),
		Port:         flex.ExpandInt64Pointer(m.Port),
		Request:      flex.ExpandStringPointer(m.Request),
		Result:       flex.ExpandStringPointer(m.Result),
		ResultCode:   flex.ExpandInt64Pointer(m.ResultCode),
		RetryDown:    flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:      flex.ExpandInt64Pointer(m.RetryUp),
		Timeout:      flex.ExpandInt64Pointer(m.Timeout),
		Transport:    flex.ExpandStringPointer(m.Transport),
		ValidateCert: flex.ExpandBoolPointer(m.ValidateCert),
	}
	return to
}

func FlattenDtcMonitorSip(ctx context.Context, from *dtc.DtcMonitorSip, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcMonitorSipAttrTypes)
	}
	m := DtcMonitorSipModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcMonitorSipAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcMonitorSipModel) Flatten(ctx context.Context, from *dtc.DtcMonitorSip, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcMonitorSipModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Ciphers = flex.FlattenStringPointer(from.Ciphers)
	m.ClientCert = flex.FlattenStringPointer(from.ClientCert)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Request = flex.FlattenStringPointer(from.Request)
	m.Result = flex.FlattenStringPointer(from.Result)
	m.ResultCode = flex.FlattenInt64Pointer(from.ResultCode)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.Transport = flex.FlattenStringPointer(from.Transport)
	m.ValidateCert = types.BoolPointerValue(from.ValidateCert)
}
