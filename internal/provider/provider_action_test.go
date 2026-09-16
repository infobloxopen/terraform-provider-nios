package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
)

func TestNIOSProviderActions(t *testing.T) {
	t.Parallel()

	actions := (&NIOSProvider{}).Actions(context.Background())
	if len(actions) != 1 {
		t.Fatalf("action factory count = %d, want 1", len(actions))
	}

	metadataResponse := action.MetadataResponse{}
	actions[0]().Metadata(context.Background(), action.MetadataRequest{ProviderTypeName: "nios"}, &metadataResponse)
	if metadataResponse.TypeName != "nios_restart_services" {
		t.Fatalf("action type name = %q, want nios_restart_services", metadataResponse.TypeName)
	}
}
