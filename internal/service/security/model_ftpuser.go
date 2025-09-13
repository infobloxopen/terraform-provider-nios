package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FtpuserModel struct {
	Ref           types.String `tfsdk:"ref"`
	CreateHomeDir types.Bool   `tfsdk:"create_home_dir"`
	ExtAttrs      types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll   types.Map    `tfsdk:"extattrs_all"`
	HomeDir       types.String `tfsdk:"home_dir"`
	Password      types.String `tfsdk:"password"`
	Permission    types.String `tfsdk:"permission"`
	Username      types.String `tfsdk:"username"`
}

var FtpuserAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
	"create_home_dir": types.BoolType,
	"extattrs":        types.MapType{ElemType: types.StringType},
	"home_dir":        types.StringType,
	"password":        types.StringType,
	"permission":      types.StringType,
	"username":        types.StringType,
}

var FtpuserResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"create_home_dir": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to create the home directory with the user name or to use the existing directory as the home directory.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"home_dir": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The absolute path of the FTP user's home directory.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FTP user password.",
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FTP user permission.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FTP user name.",
	},
}

func ExpandFtpuser(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Ftpuser {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FtpuserModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FtpuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Ftpuser {
	if m == nil {
		return nil
	}
	to := &security.Ftpuser{
		Ref:           flex.ExpandStringPointer(m.Ref),
		CreateHomeDir: flex.ExpandBoolPointer(m.CreateHomeDir),
		ExtAttrs:      ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		HomeDir:       flex.ExpandStringPointer(m.HomeDir),
		Password:      flex.ExpandStringPointer(m.Password),
		Permission:    flex.ExpandStringPointer(m.Permission),
		Username:      flex.ExpandStringPointer(m.Username),
	}
	return to
}

func FlattenFtpuser(ctx context.Context, from *security.Ftpuser, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FtpuserAttrTypes)
	}
	m := FtpuserModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FtpuserAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FtpuserModel) Flatten(ctx context.Context, from *security.Ftpuser, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FtpuserModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CreateHomeDir = types.BoolPointerValue(from.CreateHomeDir)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.HomeDir = flex.FlattenStringPointer(from.HomeDir)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.Username = flex.FlattenStringPointer(from.Username)
}
