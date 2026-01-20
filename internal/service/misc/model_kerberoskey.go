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

type KerberoskeyModel struct {
	Ref             types.String `tfsdk:"ref"`
	Uuid            types.String `tfsdk:"uuid"`
	Domain          types.String `tfsdk:"domain"`
	Enctype         types.String `tfsdk:"enctype"`
	InUse           types.Bool   `tfsdk:"in_use"`
	Members         types.List   `tfsdk:"members"`
	Principal       types.String `tfsdk:"principal"`
	UploadTimestamp types.Int64  `tfsdk:"upload_timestamp"`
	Version         types.Int64  `tfsdk:"version"`
}

var KerberoskeyAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"uuid":             types.StringType,
	"domain":           types.StringType,
	"enctype":          types.StringType,
	"in_use":           types.BoolType,
	"members":          types.ListType{ElemType: types.StringType},
	"principal":        types.StringType,
	"upload_timestamp": types.Int64Type,
	"version":          types.Int64Type,
}

var KerberoskeyResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Kerberos domain name.",
	},
	"enctype": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Kerberos key encryption type.",
	},
	"in_use": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the Kerberos key is assigned to the Grid or Grid member.",
	},
	"members": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of hostnames and services of Grid members where the key is assigned or Grid/DHCP4 or Grid/DHCP6 or Grid/DNS.",
	},
	"principal": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The principal of the Kerberos key object.",
	},
	"upload_timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp of the Kerberos key upload operation.",
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Kerberos key version number (KVNO).",
	},
}

func ExpandKerberoskey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Kerberoskey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KerberoskeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KerberoskeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Kerberoskey {
	if m == nil {
		return nil
	}
	to := &misc.Kerberoskey{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenKerberoskey(ctx context.Context, from *misc.Kerberoskey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KerberoskeyAttrTypes)
	}
	m := KerberoskeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KerberoskeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KerberoskeyModel) Flatten(ctx context.Context, from *misc.Kerberoskey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KerberoskeyModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Enctype = flex.FlattenStringPointer(from.Enctype)
	m.InUse = types.BoolPointerValue(from.InUse)
	m.Members = flex.FlattenFrameworkListString(ctx, from.Members, diags)
	m.Principal = flex.FlattenStringPointer(from.Principal)
	m.UploadTimestamp = flex.FlattenInt64Pointer(from.UploadTimestamp)
	m.Version = flex.FlattenInt64Pointer(from.Version)
}
