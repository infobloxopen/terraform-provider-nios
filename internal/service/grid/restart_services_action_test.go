package grid

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
)

func TestRestartServicesActionMetadataAndSchema(t *testing.T) {
	t.Parallel()

	actionUnderTest := NewRestartServicesAction()
	metadataResponse := action.MetadataResponse{}
	actionUnderTest.Metadata(context.Background(), action.MetadataRequest{ProviderTypeName: "nios"}, &metadataResponse)
	if metadataResponse.TypeName != "nios_restart_services" {
		t.Fatalf("action type name = %q, want nios_restart_services", metadataResponse.TypeName)
	}

	schemaResponse := action.SchemaResponse{}
	actionUnderTest.Schema(context.Background(), action.SchemaRequest{}, &schemaResponse)
	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", schemaResponse.Diagnostics)
	}
	if diagnostics := schemaResponse.Schema.ValidateImplementation(context.Background()); diagnostics.HasError() {
		t.Fatalf("schema implementation diagnostics: %v", diagnostics)
	}
	if len(schemaResponse.Schema.Attributes) != 6 {
		t.Fatalf("schema attribute count = %d, want 6", len(schemaResponse.Schema.Attributes))
	}
}

func TestRestartServicesActionConfigureRejectsUnexpectedType(t *testing.T) {
	t.Parallel()

	actionUnderTest := &RestartServicesAction{}
	response := action.ConfigureResponse{}
	actionUnderTest.Configure(context.Background(), action.ConfigureRequest{ProviderData: "not a client"}, &response)

	if !response.Diagnostics.HasError() {
		t.Fatal("Configure() returned no error for an unexpected provider data type")
	}
	if got := response.Diagnostics[0].Summary(); got != "Unexpected Action Configure Type" {
		t.Fatalf("diagnostic summary = %q, want Unexpected Action Configure Type", got)
	}
}

func TestRestartServicesActionInvokeUsesDefaults(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"result":[{"_ref":"grid/b25lLmNsdXN0ZXIkMA:Infoblox"}]}`))
			return
		}

		var got restartServicesRequest
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode restart request: %v", err)
		}
		want := restartServicesRequest{
			RestartOption: "RESTART_IF_NEEDED",
			Services:      []string{"ALL"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("restart request = %#v, want %#v", got, want)
		}
		_, _ = w.Write([]byte(`null`))
	}))
	defer server.Close()

	client := niosclient.NewAPIClient(
		option.WithNIOSHostUrl(server.URL),
		option.WithNIOSUsername("admin"),
		option.WithNIOSPassword("secret"),
		option.WithDebug(false),
	)
	actionUnderTest := &RestartServicesAction{client: client}
	schemaResponse := action.SchemaResponse{}
	actionUnderTest.Schema(context.Background(), action.SchemaRequest{}, &schemaResponse)
	listOfStrings := tftypes.List{ElementType: tftypes.String}
	rawConfig := tftypes.NewValue(schemaResponse.Schema.Type().TerraformType(context.Background()), map[string]tftypes.Value{
		"groups":         tftypes.NewValue(listOfStrings, nil),
		"members":        tftypes.NewValue(listOfStrings, nil),
		"mode":           tftypes.NewValue(tftypes.String, nil),
		"restart_option": tftypes.NewValue(tftypes.String, nil),
		"services":       tftypes.NewValue(listOfStrings, nil),
		"user_name":      tftypes.NewValue(tftypes.String, nil),
	})
	response := action.InvokeResponse{}
	actionUnderTest.Invoke(context.Background(), action.InvokeRequest{
		Config: tfsdk.Config{Raw: rawConfig, Schema: schemaResponse.Schema},
	}, &response)

	if response.Diagnostics.HasError() {
		t.Fatalf("Invoke() diagnostics: %v", response.Diagnostics)
	}
}

func TestInvokeRestartServices(t *testing.T) {
	t.Parallel()

	var requestCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++

		switch requestCount {
		case 1:
			if r.Method != http.MethodGet {
				t.Errorf("grid lookup method = %q, want GET", r.Method)
			}
			if r.URL.Path != "/wapi/v2.13.6/grid" {
				t.Errorf("grid lookup path = %q, want /wapi/v2.13.6/grid", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"result":[{"_ref":"grid/b25lLmNsdXN0ZXIkMA:Infoblox"}]}`))
		case 2:
			if r.Method != http.MethodPost {
				t.Errorf("restart method = %q, want POST", r.Method)
			}
			if r.URL.Path != "/wapi/v2.13.6/grid/b25lLmNsdXN0ZXIkMA:Infoblox" {
				t.Errorf("restart path = %q, want grid reference path", r.URL.Path)
			}
			if got := r.URL.Query().Get("_function"); got != "restartservices" {
				t.Errorf("_function = %q, want restartservices", got)
			}

			var got restartServicesRequest
			if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
				t.Fatalf("decode restart request: %v", err)
			}
			want := restartServicesRequest{
				Members:       []string{"member.example.org"},
				Mode:          "SEQUENTIAL",
				RestartOption: "RESTART_IF_NEEDED",
				Services:      []string{"DHCP"},
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("restart request = %#v, want %#v", got, want)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`null`))
		default:
			t.Errorf("unexpected request %d: %s %s", requestCount, r.Method, r.URL)
		}
	}))
	defer server.Close()

	client := niosclient.NewAPIClient(
		option.WithNIOSHostUrl(server.URL),
		option.WithNIOSUsername("admin"),
		option.WithNIOSPassword("secret"),
		option.WithDebug(false),
	)
	payload := restartServicesRequest{
		Members:       []string{"member.example.org"},
		Mode:          "SEQUENTIAL",
		RestartOption: "RESTART_IF_NEEDED",
		Services:      []string{"DHCP"},
	}

	if err := invokeRestartServices(context.Background(), client, payload); err != nil {
		t.Fatalf("invokeRestartServices() error = %v", err)
	}
	if requestCount != 2 {
		t.Fatalf("request count = %d, want 2", requestCount)
	}
}
