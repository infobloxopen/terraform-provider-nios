package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
)

type TftpfiledirModel struct {
	Ref             types.String `tfsdk:"ref"`
	Uuid        types.String `tfsdk:"uuid"`
	Directory       types.String `tfsdk:"directory"`
	IsSyncedToGm    types.Bool   `tfsdk:"is_synced_to_gm"`
	LastModify      types.Int64  `tfsdk:"last_modify"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	VtftpDirMembers types.List   `tfsdk:"vtftp_dir_members"`
}

var TftpfiledirAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"uuid":        types.StringType,
	"directory":         types.StringType,
	"is_synced_to_gm":   types.BoolType,
	"last_modify":       types.Int64Type,
	"name":              types.StringType,
	"type":              types.StringType,
	"vtftp_dir_members": types.ListType{ElemType: types.ObjectType{AttrTypes: TftpfiledirVtftpDirMembersAttrTypes}},
}

var TftpfiledirResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Universally Unique ID assigned for this object.",
	},
	"directory": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("/"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The path to the directory that contains file or subdirectory.",
	},
	"is_synced_to_gm": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the TFTP entity is synchronized to Grid Master.",
	},
	"last_modify": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the file or directory was last modified.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The TFTP directory or file name.",
	},
	"type": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		Validators: []validator.String{
			stringvalidator.OneOf("DIRECTORY", "FILE"),
		},
		MarkdownDescription: "The type of TFTP file system entity (directory or file). TYPE `FILE` is not supported through terraform provider and is reserved for future use.",
	},
	"vtftp_dir_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TftpfiledirVtftpDirMembersResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The replication members with TFTP client addresses where this virtual folder is applicable.",
	},
}

func (m *TftpfiledirModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *misc.Tftpfiledir {
	if m == nil {
		return nil
	}
	to := &misc.Tftpfiledir{
		Name: flex.ExpandStringPointer(m.Name),
	}
	if isCreate {
		to.Directory = flex.ExpandStringPointer(m.Directory)
		to.Type = flex.ExpandStringPointer(m.Type)
	}
	if !isCreate {
		to.VtftpDirMembers = flex.ExpandFrameworkListNestedBlock(ctx, m.VtftpDirMembers, diags, ExpandTftpfiledirVtftpDirMembers)
	}
	return to
}

func FlattenTftpfiledir(ctx context.Context, from *misc.Tftpfiledir, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TftpfiledirAttrTypes)
	}
	m := TftpfiledirModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TftpfiledirAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TftpfiledirModel) Flatten(ctx context.Context, from *misc.Tftpfiledir, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TftpfiledirModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Directory = flex.FlattenStringPointer(from.Directory)
	m.IsSyncedToGm = types.BoolPointerValue(from.IsSyncedToGm)
	m.LastModify = flex.FlattenInt64Pointer(from.LastModify)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.VtftpDirMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.VtftpDirMembers, TftpfiledirVtftpDirMembersAttrTypes, diags, FlattenTftpfiledirVtftpDirMembers)
}
