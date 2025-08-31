package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type FixedaddressCliCredentialsModel struct {
	User            types.String `tfsdk:"user"`
	Password        types.String `tfsdk:"password"`
	CredentialType  types.String `tfsdk:"credential_type"`
	Comment         types.String `tfsdk:"comment"`
	Id              types.Int64  `tfsdk:"id"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

var FixedaddressCliCredentialsAttrTypes = map[string]attr.Type{
	"user":             types.StringType,
	"password":         types.StringType,
	"credential_type":  types.StringType,
	"comment":          types.StringType,
	"id":               types.Int64Type,
	"credential_group": types.StringType,
}

var FixedaddressCliCredentialsResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The CLI user name.",
	},
	"password": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The CLI password.",
	},
	"credential_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ENABLE_SSH", "ENABLE_TELNET", "SSH", "TELNET"),
		},
		MarkdownDescription: "The type of the credential.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for the credential.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Credentials ID.",
	},
	"credential_group": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Group for the CLI credential.",
	},
}

func ExpandFixedaddressCliCredentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressCliCredentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressCliCredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressCliCredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressCliCredentials {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressCliCredentials{
		User:            flex.ExpandStringPointer(m.User),
		Password:        flex.ExpandStringPointer(m.Password),
		CredentialType:  flex.ExpandStringPointer(m.CredentialType),
		Comment:         flex.ExpandStringPointer(m.Comment),
		CredentialGroup: flex.ExpandStringPointer(m.CredentialGroup),
	}
	return to
}

func FlattenFixedaddressCliCredentials(ctx context.Context, from *dhcp.FixedaddressCliCredentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressCliCredentialsAttrTypes)
	}
	m := FixedaddressCliCredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressCliCredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressCliCredentialsModel) Flatten(ctx context.Context, from *dhcp.FixedaddressCliCredentials, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressCliCredentialsModel{}
	}
	m.User = flex.FlattenStringPointer(from.User)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.CredentialType = flex.FlattenStringPointer(from.CredentialType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
}
