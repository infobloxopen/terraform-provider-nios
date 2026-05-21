package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordHostCliCredentialsModel struct {
	User            types.String `tfsdk:"user"`
	Password        types.String `tfsdk:"password"`
	CredentialType  types.String `tfsdk:"credential_type"`
	Comment         types.String `tfsdk:"comment"`
	Id              types.Int64  `tfsdk:"id"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

var RecordHostCliCredentialsAttrTypes = map[string]attr.Type{
	"user":             types.StringType,
	"password":         types.StringType,
	"credential_type":  types.StringType,
	"comment":          types.StringType,
	"id":               types.Int64Type,
	"credential_group": types.StringType,
}

var RecordHostCliCredentialsResourceSchemaAttributes = map[string]schema.Attribute{
	"user": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The CLI user name.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Sensitive:           true,
		MarkdownDescription: "The CLI password.",
	},
	"credential_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of the credential.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The commment for the credential.",
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

func ExpandRecordHostCliCredentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHostCliCredentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostCliCredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordHostCliCredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHostCliCredentials {
	if m == nil {
		return nil
	}
	to := &dns.RecordHostCliCredentials{
		User:            flex.ExpandStringPointer(m.User),
		Password:        flex.ExpandStringPointer(m.Password),
		CredentialType:  flex.ExpandStringPointer(m.CredentialType),
		Comment:         flex.ExpandStringPointer(m.Comment),
		CredentialGroup: flex.ExpandStringPointer(m.CredentialGroup),
	}
	return to
}

func FlattenRecordHostCliCredentials(ctx context.Context, from *dns.RecordHostCliCredentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostCliCredentialsAttrTypes)
	}
	m := RecordHostCliCredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostCliCredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordHostCliCredentialsModel) Flatten(ctx context.Context, from *dns.RecordHostCliCredentials, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordHostCliCredentialsModel{}
	}
	m.User = flex.FlattenStringPointer(from.User)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.CredentialType = flex.FlattenStringPointer(from.CredentialType)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
}

func preserveCliCredentialPasswords(ctx context.Context, saved types.List, current types.List, diags *diag.Diagnostics) types.List {
	if saved.IsNull() || saved.IsUnknown() || current.IsNull() || current.IsUnknown() {
		return current
	}

	var savedModels []RecordHostCliCredentialsModel
	diags.Append(saved.ElementsAs(ctx, &savedModels, false)...)
	if diags.HasError() {
		return current
	}

	var currentModels []RecordHostCliCredentialsModel
	diags.Append(current.ElementsAs(ctx, &currentModels, false)...)
	if diags.HasError() {
		return current
	}

	for i := range currentModels {
		if i < len(savedModels) && (currentModels[i].Password.IsNull() || currentModels[i].Password.ValueString() == "") {
			currentModels[i].Password = savedModels[i].Password
		}
	}

	result, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: RecordHostCliCredentialsAttrTypes}, currentModels)
	diags.Append(d...)
	return result
}
