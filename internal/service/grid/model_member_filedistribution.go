package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberFiledistributionModel struct {
	Ref               types.String `tfsdk:"ref"`
	Uuid              types.String `tfsdk:"uuid"`
	AllowUploads      types.Bool   `tfsdk:"allow_uploads"`
	Comment           types.String `tfsdk:"comment"`
	EnableFtp         types.Bool   `tfsdk:"enable_ftp"`
	EnableFtpFilelist types.Bool   `tfsdk:"enable_ftp_filelist"`
	EnableFtpPassive  types.Bool   `tfsdk:"enable_ftp_passive"`
	EnableHttp        types.Bool   `tfsdk:"enable_http"`
	EnableHttpAcl     types.Bool   `tfsdk:"enable_http_acl"`
	EnableTftp        types.Bool   `tfsdk:"enable_tftp"`
	FtpAcls           types.List   `tfsdk:"ftp_acls"`
	FtpPort           types.Int64  `tfsdk:"ftp_port"`
	FtpStatus         types.String `tfsdk:"ftp_status"`
	HostName          types.String `tfsdk:"host_name"`
	HttpAcls          types.List   `tfsdk:"http_acls"`
	HttpStatus        types.String `tfsdk:"http_status"`
	Ipv4Address       types.String `tfsdk:"ipv4_address"`
	Ipv6Address       types.String `tfsdk:"ipv6_address"`
	Status            types.String `tfsdk:"status"`
	TftpAcls          types.List   `tfsdk:"tftp_acls"`
	TftpPort          types.Int64  `tfsdk:"tftp_port"`
	TftpStatus        types.String `tfsdk:"tftp_status"`
	UseAllowUploads   types.Bool   `tfsdk:"use_allow_uploads"`
}

var MemberFiledistributionAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"uuid":                types.StringType,
	"allow_uploads":       types.BoolType,
	"comment":             types.StringType,
	"enable_ftp":          types.BoolType,
	"enable_ftp_filelist": types.BoolType,
	"enable_ftp_passive":  types.BoolType,
	"enable_http":         types.BoolType,
	"enable_http_acl":     types.BoolType,
	"enable_tftp":         types.BoolType,
	"ftp_acls":            types.ListType{ElemType: types.ObjectType{AttrTypes: MemberFiledistributionFtpAclsAttrTypes}},
	"ftp_port":            types.Int64Type,
	"ftp_status":          types.StringType,
	"host_name":           types.StringType,
	"http_acls":           types.ListType{ElemType: types.ObjectType{AttrTypes: MemberFiledistributionHttpAclsAttrTypes}},
	"http_status":         types.StringType,
	"ipv4_address":        types.StringType,
	"ipv6_address":        types.StringType,
	"status":              types.StringType,
	"tftp_acls":           types.ListType{ElemType: types.ObjectType{AttrTypes: MemberFiledistributionTftpAclsAttrTypes}},
	"tftp_port":           types.Int64Type,
	"tftp_status":         types.StringType,
	"use_allow_uploads":   types.BoolType,
}

var MemberFiledistributionResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"allow_uploads": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether uploads to the Grid member are allowed.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member descriptive comment.",
	},
	"enable_ftp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the FTP prtocol is enabled for file distribution.",
	},
	"enable_ftp_filelist": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the LIST command for FTP is enabled.",
	},
	"enable_ftp_passive": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the passive mode for FTP is enabled.",
	},
	"enable_http": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the HTTP prtocol is enabled for file distribution.",
	},
	"enable_http_acl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the HTTP prtocol access control (AC) settings are enabled.",
	},
	"enable_tftp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the TFTP prtocol is enabled for file distribution.",
	},
	"ftp_acls": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberFiledistributionFtpAclsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Access control (AC) settings for the FTP protocol.",
	},
	"ftp_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The network port used by the FTP protocol.",
	},
	"ftp_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The FTP protocol status.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member host name.",
	},
	"http_acls": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberFiledistributionHttpAclsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Access control (AC) settings for the HTTP protocol.",
	},
	"http_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The HTTP protocol status.",
	},
	"ipv4_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 address of the Grid member.",
	},
	"ipv6_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 address of the Grid member.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member file distribution status.",
	},
	"tftp_acls": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberFiledistributionTftpAclsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The access control (AC) settings for the TFTP protocol.",
	},
	"tftp_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The network port used by the TFTP protocol.",
	},
	"tftp_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TFTP protocol status.",
	},
	"use_allow_uploads": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: allow_uploads",
	},
}

func ExpandMemberFiledistribution(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberFiledistribution {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberFiledistributionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberFiledistributionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberFiledistribution {
	if m == nil {
		return nil
	}
	to := &grid.MemberFiledistribution{
		Ref:               flex.ExpandStringPointer(m.Ref),
		Uuid:              flex.ExpandStringPointer(m.Uuid),
		AllowUploads:      flex.ExpandBoolPointer(m.AllowUploads),
		EnableFtp:         flex.ExpandBoolPointer(m.EnableFtp),
		EnableFtpFilelist: flex.ExpandBoolPointer(m.EnableFtpFilelist),
		EnableFtpPassive:  flex.ExpandBoolPointer(m.EnableFtpPassive),
		EnableHttp:        flex.ExpandBoolPointer(m.EnableHttp),
		EnableHttpAcl:     flex.ExpandBoolPointer(m.EnableHttpAcl),
		EnableTftp:        flex.ExpandBoolPointer(m.EnableTftp),
		FtpAcls:           flex.ExpandFrameworkListNestedBlock(ctx, m.FtpAcls, diags, ExpandMemberFiledistributionFtpAcls),
		FtpPort:           flex.ExpandInt64Pointer(m.FtpPort),
		HttpAcls:          flex.ExpandFrameworkListNestedBlock(ctx, m.HttpAcls, diags, ExpandMemberFiledistributionHttpAcls),
		TftpAcls:          flex.ExpandFrameworkListNestedBlock(ctx, m.TftpAcls, diags, ExpandMemberFiledistributionTftpAcls),
		TftpPort:          flex.ExpandInt64Pointer(m.TftpPort),
		UseAllowUploads:   flex.ExpandBoolPointer(m.UseAllowUploads),
	}
	return to
}

func FlattenMemberFiledistribution(ctx context.Context, from *grid.MemberFiledistribution, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberFiledistributionAttrTypes)
	}
	m := MemberFiledistributionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberFiledistributionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberFiledistributionModel) Flatten(ctx context.Context, from *grid.MemberFiledistribution, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberFiledistributionModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowUploads = types.BoolPointerValue(from.AllowUploads)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EnableFtp = types.BoolPointerValue(from.EnableFtp)
	m.EnableFtpFilelist = types.BoolPointerValue(from.EnableFtpFilelist)
	m.EnableFtpPassive = types.BoolPointerValue(from.EnableFtpPassive)
	m.EnableHttp = types.BoolPointerValue(from.EnableHttp)
	m.EnableHttpAcl = types.BoolPointerValue(from.EnableHttpAcl)
	m.EnableTftp = types.BoolPointerValue(from.EnableTftp)
	m.FtpAcls = flex.FlattenFrameworkListNestedBlock(ctx, from.FtpAcls, MemberFiledistributionFtpAclsAttrTypes, diags, FlattenMemberFiledistributionFtpAcls)
	m.FtpPort = flex.FlattenInt64Pointer(from.FtpPort)
	m.FtpStatus = flex.FlattenStringPointer(from.FtpStatus)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.HttpAcls = flex.FlattenFrameworkListNestedBlock(ctx, from.HttpAcls, MemberFiledistributionHttpAclsAttrTypes, diags, FlattenMemberFiledistributionHttpAcls)
	m.HttpStatus = flex.FlattenStringPointer(from.HttpStatus)
	m.Ipv4Address = flex.FlattenStringPointer(from.Ipv4Address)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.TftpAcls = flex.FlattenFrameworkListNestedBlock(ctx, from.TftpAcls, MemberFiledistributionTftpAclsAttrTypes, diags, FlattenMemberFiledistributionTftpAcls)
	m.TftpPort = flex.FlattenInt64Pointer(from.TftpPort)
	m.TftpStatus = flex.FlattenStringPointer(from.TftpStatus)
	m.UseAllowUploads = types.BoolPointerValue(from.UseAllowUploads)
}
