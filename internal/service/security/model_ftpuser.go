package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
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
	"extattrs_all":    types.MapType{ElemType: types.StringType},
	"home_dir":        types.StringType,
	"password":        types.StringType,
	"permission":      types.StringType,
	"username":        types.StringType,
}

var FtpuserResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"create_home_dir": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether to create the home directory with the user name or to use the existing directory as the home directory. Default is true. Cannot be used together with 'home_dir'.",
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
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
	},
	"home_dir": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The absolute path of the FTP user's home directory. set create_home_dir to false to set home_dir",
	},
	"password": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The FTP user password.",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("RO"),
		Validators: []validator.String{
			stringvalidator.OneOf("RO", "RW"),
		},
		MarkdownDescription: "The FTP user permission.",
	},
	"username": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The FTP user name.",
	},
}

func (m *FtpuserModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *security.Ftpuser {
	if m == nil {
		return nil
	}
	to := &security.Ftpuser{
		ExtAttrs:   ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	if isCreate {
		to.CreateHomeDir = flex.ExpandBoolPointer(m.CreateHomeDir)
		to.HomeDir = flex.ExpandStringPointer(m.HomeDir)
		to.Password = flex.ExpandStringPointer(m.Password)
		to.Username = flex.ExpandStringPointer(m.Username)
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
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.HomeDir = flex.FlattenStringPointer(from.HomeDir)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.Username = flex.FlattenStringPointer(from.Username)
}
