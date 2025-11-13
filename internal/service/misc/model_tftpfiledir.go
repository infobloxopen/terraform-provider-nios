package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type TftpfiledirModel struct {
	Ref             types.String `tfsdk:"ref"`
	Directory       types.String `tfsdk:"directory"`
	IsSyncedToGm    types.Bool   `tfsdk:"is_synced_to_gm"`
	LastModify      types.Int64  `tfsdk:"last_modify"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	VtftpDirMembers types.List   `tfsdk:"vtftp_dir_members"`
}

var TftpfiledirAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"directory":         types.StringType,
	"is_synced_to_gm":   types.BoolType,
	"last_modify":       types.Int64Type,
	"name":              types.StringType,
	"type":              types.StringType,
	"vtftp_dir_members": types.ListType{ElemType: types.ObjectType{AttrTypes: TftpfiledirVtftpDirMembersAttrTypes}},
}

var TftpfiledirResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"directory": schema.StringAttribute{
		Optional:            true,
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
		Optional:            true,
		MarkdownDescription: "The TFTP directory or file name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of TFTP file system entity (directory or file).",
	},
	"vtftp_dir_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TftpfiledirVtftpDirMembersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The replication members with TFTP client addresses where this virtual folder is applicable.",
	},
}

func ExpandTftpfiledir(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Tftpfiledir {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TftpfiledirModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TftpfiledirModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Tftpfiledir {
	if m == nil {
		return nil
	}
	to := &misc.Tftpfiledir{
		Ref:             flex.ExpandStringPointer(m.Ref),
		Directory:       flex.ExpandStringPointer(m.Directory),
		Name:            flex.ExpandStringPointer(m.Name),
		Type:            flex.ExpandStringPointer(m.Type),
		VtftpDirMembers: flex.ExpandFrameworkListNestedBlock(ctx, m.VtftpDirMembers, diags, ExpandTftpfiledirVtftpDirMembers),
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
	m.Directory = flex.FlattenStringPointer(from.Directory)
	m.IsSyncedToGm = types.BoolPointerValue(from.IsSyncedToGm)
	m.LastModify = flex.FlattenInt64Pointer(from.LastModify)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.VtftpDirMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.VtftpDirMembers, TftpfiledirVtftpDirMembersAttrTypes, diags, FlattenTftpfiledirVtftpDirMembers)
}
