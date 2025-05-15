package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyz"
	defaultKey   = "managed_by"
	defaultValue = "terraform"
)

var (
	// NIOSClient will be used to do verification tests
	NIOSClient *niosclient.APIClient

	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"nios": providerserver.NewProtocol6WithError(provider.New("test", "test")()),
	}
)

// RandomNameWithPrefix generates a random name with the given prefix.
// This is used in the acceptance tests where a unique name is required for the resource.
func RandomNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, RandomName())
}

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func RandomName() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PreCheck(t *testing.T) {
	hostURL := os.Getenv("NIOS_HOST_URL")
	if hostURL == "" {
		t.Fatal("NIOS_HOST_URL must be set for acceptance tests")
	}

	auth := os.Getenv("NIOS_AUTH")
	if auth == "" {
		t.Fatal("NIOS_AUTH must be set for acceptance tests")
	}

	NIOSClient = niosclient.NewAPIClient(
		option.WithClientName("terraform-acceptance-tests"),
		option.WithNIOSHostUrl(hostURL),
		option.WithNIOSAuth(auth),
		option.WithDebug(true),
	)
}
