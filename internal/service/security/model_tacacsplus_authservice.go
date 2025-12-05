package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type TacacsplusAuthserviceModel struct {
	Ref         types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	AcctRetries types.Int64  `tfsdk:"acct_retries"`
	AcctTimeout types.Int64  `tfsdk:"acct_timeout"`
	AuthRetries types.Int64  `tfsdk:"auth_retries"`
	AuthTimeout types.Int64  `tfsdk:"auth_timeout"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	Name        types.String `tfsdk:"name"`
	Servers     types.List   `tfsdk:"servers"`
}

var TacacsplusAuthserviceAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
    "uuid":        types.StringType,
	"acct_retries": types.Int64Type,
	"acct_timeout": types.Int64Type,
	"auth_retries": types.Int64Type,
	"auth_timeout": types.Int64Type,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"name":         types.StringType,
	"servers":      types.ListType{ElemType: types.ObjectType{AttrTypes: TacacsplusAuthserviceServersAttrTypes}},
}

var TacacsplusAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"acct_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of the accounting retries before giving up and moving on to the next server.",
	},
	"acct_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The accounting retry period in milliseconds.",
	},
	"auth_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of the authentication/authorization retries before giving up and moving on to the next server.",
	},
	"auth_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The authentication/authorization timeout period in milliseconds.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TACACS+ authentication service descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the TACACS+ authentication service object is disabled.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TACACS+ authentication service name.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TacacsplusAuthserviceServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of the TACACS+ servers used for authentication.",
	},
}

func ExpandTacacsplusAuthservice(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.TacacsplusAuthservice {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TacacsplusAuthserviceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TacacsplusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.TacacsplusAuthservice {
	if m == nil {
		return nil
	}
	to := &security.TacacsplusAuthservice{
		Ref:         flex.ExpandStringPointer(m.Ref),
		AcctRetries: flex.ExpandInt64Pointer(m.AcctRetries),
		AcctTimeout: flex.ExpandInt64Pointer(m.AcctTimeout),
		AuthRetries: flex.ExpandInt64Pointer(m.AuthRetries),
		AuthTimeout: flex.ExpandInt64Pointer(m.AuthTimeout),
		Comment:     flex.ExpandStringPointer(m.Comment),
		Disable:     flex.ExpandBoolPointer(m.Disable),
		Name:        flex.ExpandStringPointer(m.Name),
		Servers:     flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandTacacsplusAuthserviceServers),
	}
	return to
}

func FlattenTacacsplusAuthservice(ctx context.Context, from *security.TacacsplusAuthservice, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TacacsplusAuthserviceAttrTypes)
	}
	m := TacacsplusAuthserviceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TacacsplusAuthserviceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TacacsplusAuthserviceModel) Flatten(ctx context.Context, from *security.TacacsplusAuthservice, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TacacsplusAuthserviceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AcctRetries = flex.FlattenInt64Pointer(from.AcctRetries)
	m.AcctTimeout = flex.FlattenInt64Pointer(from.AcctTimeout)
	m.AuthRetries = flex.FlattenInt64Pointer(from.AuthRetries)
	m.AuthTimeout = flex.FlattenInt64Pointer(from.AuthTimeout)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, TacacsplusAuthserviceServersAttrTypes, diags, FlattenTacacsplusAuthserviceServers)
}
