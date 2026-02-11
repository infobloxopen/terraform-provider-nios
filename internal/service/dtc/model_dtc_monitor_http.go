package dtc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type DtcMonitorHttpModel struct {
	Ref                 types.String `tfsdk:"ref"`
	Uuid                types.String `tfsdk:"uuid"`
	Ciphers             types.String `tfsdk:"ciphers"`
	ClientCert          types.String `tfsdk:"client_cert"`
	Comment             types.String `tfsdk:"comment"`
	ContentCheck        types.String `tfsdk:"content_check"`
	ContentCheckInput   types.String `tfsdk:"content_check_input"`
	ContentCheckOp      types.String `tfsdk:"content_check_op"`
	ContentCheckRegex   types.String `tfsdk:"content_check_regex"`
	ContentExtractGroup types.Int64  `tfsdk:"content_extract_group"`
	ContentExtractType  types.String `tfsdk:"content_extract_type"`
	ContentExtractValue types.String `tfsdk:"content_extract_value"`
	EnableSni           types.Bool   `tfsdk:"enable_sni"`
	ExtAttrs            types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll         types.Map    `tfsdk:"extattrs_all"`
	Interval            types.Int64  `tfsdk:"interval"`
	Name                types.String `tfsdk:"name"`
	Port                types.Int64  `tfsdk:"port"`
	Request             types.String `tfsdk:"request"`
	Result              types.String `tfsdk:"result"`
	ResultCode          types.Int64  `tfsdk:"result_code"`
	RetryDown           types.Int64  `tfsdk:"retry_down"`
	RetryUp             types.Int64  `tfsdk:"retry_up"`
	Secure              types.Bool   `tfsdk:"secure"`
	Timeout             types.Int64  `tfsdk:"timeout"`
	ValidateCert        types.Bool   `tfsdk:"validate_cert"`
}

var DtcMonitorHttpAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"uuid":                  types.StringType,
	"ciphers":               types.StringType,
	"client_cert":           types.StringType,
	"comment":               types.StringType,
	"content_check":         types.StringType,
	"content_check_input":   types.StringType,
	"content_check_op":      types.StringType,
	"content_check_regex":   types.StringType,
	"content_extract_group": types.Int64Type,
	"content_extract_type":  types.StringType,
	"content_extract_value": types.StringType,
	"enable_sni":            types.BoolType,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"extattrs_all":          types.MapType{ElemType: types.StringType},
	"interval":              types.Int64Type,
	"name":                  types.StringType,
	"port":                  types.Int64Type,
	"request":               types.StringType,
	"result":                types.StringType,
	"result_code":           types.Int64Type,
	"retry_down":            types.Int64Type,
	"retry_up":              types.Int64Type,
	"secure":                types.BoolType,
	"timeout":               types.Int64Type,
	"validate_cert":         types.BoolType,
}

var DtcMonitorHttpResourceSchemaAttributes = map[string]schema.Attribute{
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
		MarkdownDescription: "An optional cipher list for a secure HTTP/S connection.",
	},
	"client_cert": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "An optional client certificate, supplied in a secure HTTP/S mode if present.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
	},
	"content_check": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "EXTRACT", "MATCH"),
		},
		MarkdownDescription: "The content check type.",
	},
	"content_check_input": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ALL"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALL", "HEADERS", "BODY"),
		},
		MarkdownDescription: "A portion of response to use as input for content check.",
	},
	"content_check_op": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("EQ", "GEQ", "LEQ", "NEQ"),
		},
		MarkdownDescription: "A content check success criteria operator.",
	},
	"content_check_regex": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A content check regular expression.",
	},
	"content_extract_group": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 8),
		},
		MarkdownDescription: "A content extraction sub-expression to extract.",
	},
	"content_extract_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("STRING"),
		Validators: []validator.String{
			stringvalidator.OneOf("STRING", "INTEGER"),
		},
		MarkdownDescription: "A content extraction expected type for the extracted data.",
	},
	"content_extract_value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A content extraction value to compare with extracted result.",
	},
	"enable_sni": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the Server Name Indication (SNI) for HTTPS monitor is enabled.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for HTTP health check.",
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
		Default:             int64default.StaticInt64(80),
		MarkdownDescription: "Port for HTTP requests.",
	},
	"request": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "An HTTP request to send.",
	},
	"result": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ANY"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "CODE_IS", "CODE_IS_NOT"),
		},
		MarkdownDescription: "The type of an expected result.",
	},
	"result_code": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(200),
		MarkdownDescription: "The expected return code.",
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
	"secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The connection security status.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for HTTP health check in seconds.",
	},
	"validate_cert": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the validation of the remote server's certificate is enabled.",
	},
}

func (m *DtcMonitorHttpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcMonitorHttp {
	if m == nil {
		return nil
	}
	to := &dtc.DtcMonitorHttp{
		Ciphers:             flex.ExpandStringPointer(m.Ciphers),
		ClientCert:          flex.ExpandStringPointer(m.ClientCert),
		Comment:             flex.ExpandStringPointer(m.Comment),
		ContentCheck:        flex.ExpandStringPointer(m.ContentCheck),
		ContentCheckInput:   flex.ExpandStringPointer(m.ContentCheckInput),
		ContentCheckOp:      flex.ExpandStringPointer(m.ContentCheckOp),
		ContentCheckRegex:   flex.ExpandStringPointer(m.ContentCheckRegex),
		ContentExtractGroup: flex.ExpandInt64Pointer(m.ContentExtractGroup),
		ContentExtractType:  flex.ExpandStringPointer(m.ContentExtractType),
		ContentExtractValue: flex.ExpandStringPointer(m.ContentExtractValue),
		EnableSni:           flex.ExpandBoolPointer(m.EnableSni),
		ExtAttrs:            ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Interval:            flex.ExpandInt64Pointer(m.Interval),
		Name:                flex.ExpandStringPointer(m.Name),
		Port:                flex.ExpandInt64Pointer(m.Port),
		Request:             flex.ExpandStringPointer(m.Request),
		Result:              flex.ExpandStringPointer(m.Result),
		ResultCode:          flex.ExpandInt64Pointer(m.ResultCode),
		RetryDown:           flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:             flex.ExpandInt64Pointer(m.RetryUp),
		Secure:              flex.ExpandBoolPointer(m.Secure),
		Timeout:             flex.ExpandInt64Pointer(m.Timeout),
		ValidateCert:        flex.ExpandBoolPointer(m.ValidateCert),
	}
	return to
}

func FlattenDtcMonitorHttp(ctx context.Context, from *dtc.DtcMonitorHttp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcMonitorHttpAttrTypes)
	}
	m := DtcMonitorHttpModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcMonitorHttpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcMonitorHttpModel) Flatten(ctx context.Context, from *dtc.DtcMonitorHttp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcMonitorHttpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Ciphers = flex.FlattenStringPointer(from.Ciphers)
	m.ClientCert = flex.FlattenStringPointer(from.ClientCert)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ContentCheck = flex.FlattenStringPointer(from.ContentCheck)
	m.ContentCheckInput = flex.FlattenStringPointer(from.ContentCheckInput)
	m.ContentCheckOp = flex.FlattenStringPointer(from.ContentCheckOp)
	m.ContentCheckRegex = flex.FlattenStringPointer(from.ContentCheckRegex)
	m.ContentExtractGroup = flex.FlattenInt64Pointer(from.ContentExtractGroup)
	m.ContentExtractType = flex.FlattenStringPointer(from.ContentExtractType)
	m.ContentExtractValue = flex.FlattenStringPointer(from.ContentExtractValue)
	m.EnableSni = types.BoolPointerValue(from.EnableSni)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Request = FlattenRequestDtcMonitorHttp(ctx, from.Request, m.Request, diags)
	m.Result = flex.FlattenStringPointer(from.Result)
	m.ResultCode = flex.FlattenInt64Pointer(from.ResultCode)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Secure = types.BoolPointerValue(from.Secure)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.ValidateCert = types.BoolPointerValue(from.ValidateCert)
}

func FlattenRequestDtcMonitorHttp(ctx context.Context, fromRequest *string, planRequest types.String, diags *diag.Diagnostics) types.String {
	if fromRequest == nil {
		return types.StringNull()
	}
	requestValue := *fromRequest

	if strings.HasPrefix(requestValue, "POST") || strings.HasPrefix(requestValue, "GET") || strings.HasPrefix(requestValue, "HEAD") {
		if strings.HasSuffix(requestValue, "\n\n") && !strings.HasSuffix(planRequest.ValueString(), "\n\n") {
			requestValue = strings.TrimSuffix(requestValue, "\n\n")
		}

		if strings.Contains(requestValue, "HTTP/1.1") && !strings.Contains(planRequest.ValueString(), "\nConnection: close") {
			requestValue = strings.ReplaceAll(requestValue, "\nConnection: close", "")
		}
	}

	return types.StringValue(requestValue)
}
