package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type CsvimporttaskModel struct {
	Ref            types.String `tfsdk:"ref"`
	Action         types.String `tfsdk:"action"`
	AdminName      types.String `tfsdk:"admin_name"`
	EndTime        types.Int64  `tfsdk:"end_time"`
	FileName       types.String `tfsdk:"file_name"`
	FileSize       types.Int64  `tfsdk:"file_size"`
	ImportId       types.Int64  `tfsdk:"import_id"`
	LinesFailed    types.Int64  `tfsdk:"lines_failed"`
	LinesProcessed types.Int64  `tfsdk:"lines_processed"`
	LinesWarning   types.Int64  `tfsdk:"lines_warning"`
	OnError        types.String `tfsdk:"on_error"`
	Operation      types.String `tfsdk:"operation"`
	Separator      types.String `tfsdk:"separator"`
	StartTime      types.Int64  `tfsdk:"start_time"`
	Status         types.String `tfsdk:"status"`
	UpdateMethod   types.String `tfsdk:"update_method"`
}

var CsvimporttaskAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
	"action":          types.StringType,
	"admin_name":      types.StringType,
	"end_time":        types.Int64Type,
	"file_name":       types.StringType,
	"file_size":       types.Int64Type,
	"import_id":       types.Int64Type,
	"lines_failed":    types.Int64Type,
	"lines_processed": types.Int64Type,
	"lines_warning":   types.Int64Type,
	"on_error":        types.StringType,
	"operation":       types.StringType,
	"separator":       types.StringType,
	"start_time":      types.Int64Type,
	"status":          types.StringType,
	"update_method":   types.StringType,
}

var CsvimporttaskResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action to execute.",
	},
	"admin_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The login name of the administrator.",
	},
	"end_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The end time of this import operation.",
	},
	"file_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the file used for the import operation.",
	},
	"file_size": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The size of the file used for the import operation.",
	},
	"import_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The ID of the current import task.",
	},
	"lines_failed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of lines that encountered an error.",
	},
	"lines_processed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of lines that have been processed.",
	},
	"lines_warning": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of lines that encountered a warning.",
	},
	"on_error": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action to take when an error is encountered.",
	},
	"operation": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The operation to execute.",
	},
	"separator": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The separator to be used for the data in the CSV file.",
	},
	"start_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The start time of the import operation.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the import operation",
	},
	"update_method": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update method to be used for the operation.",
	},
}

func ExpandCsvimporttask(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Csvimporttask {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CsvimporttaskModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CsvimporttaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Csvimporttask {
	if m == nil {
		return nil
	}
	to := &misc.Csvimporttask{
		Ref:          flex.ExpandStringPointer(m.Ref),
		Action:       flex.ExpandStringPointer(m.Action),
		OnError:      flex.ExpandStringPointer(m.OnError),
		Operation:    flex.ExpandStringPointer(m.Operation),
		UpdateMethod: flex.ExpandStringPointer(m.UpdateMethod),
	}
	return to
}

func FlattenCsvimporttask(ctx context.Context, from *misc.Csvimporttask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CsvimporttaskAttrTypes)
	}
	m := CsvimporttaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CsvimporttaskAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CsvimporttaskModel) Flatten(ctx context.Context, from *misc.Csvimporttask, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CsvimporttaskModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Action = flex.FlattenStringPointer(from.Action)
	m.AdminName = flex.FlattenStringPointer(from.AdminName)
	m.EndTime = flex.FlattenInt64Pointer(from.EndTime)
	m.FileName = flex.FlattenStringPointer(from.FileName)
	m.FileSize = flex.FlattenInt64Pointer(from.FileSize)
	m.ImportId = flex.FlattenInt64Pointer(from.ImportId)
	m.LinesFailed = flex.FlattenInt64Pointer(from.LinesFailed)
	m.LinesProcessed = flex.FlattenInt64Pointer(from.LinesProcessed)
	m.LinesWarning = flex.FlattenInt64Pointer(from.LinesWarning)
	m.OnError = flex.FlattenStringPointer(from.OnError)
	m.Operation = flex.FlattenStringPointer(from.Operation)
	m.Separator = flex.FlattenStringPointer(from.Separator)
	m.StartTime = flex.FlattenInt64Pointer(from.StartTime)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.UpdateMethod = flex.FlattenStringPointer(from.UpdateMethod)
}
