package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
	"github.com/infobloxopen/terraform-provider-nios/internal/provider"
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

// RandomCIDRNetwork generates a random network with specific CIDR
func RandomCIDRNetwork() string {
	// Generate test-suitable private networks
	base := 10 + rand.Intn(246) // 10-255 for first octet
	second := rand.Intn(256)    // 0-255 for second octet
	cidr := 16 + rand.Intn(9)   // /16 to /24 (common for network containers)

	return fmt.Sprintf("%d.%d.0.0/%d", base, second, cidr)
}

func PreCheck(t *testing.T) {
	hostURL := os.Getenv("NIOS_HOST_URL")
	if hostURL == "" {
		t.Fatal("NIOS_HOST_URL must be set for acceptance tests")
	}

	username := os.Getenv("NIOS_USERNAME")
	if username == "" {
		t.Fatal("NIOS_USERNAME must be set for acceptance tests")
	}

	password := os.Getenv("NIOS_PASSWORD")
	if password == "" {
		t.Fatal("NIOS_PASSWORD must be set for acceptance tests")
	}

	NIOSClient = niosclient.NewAPIClient(
		option.WithClientName("terraform-acceptance-tests"),
		option.WithNIOSHostUrl(hostURL),
		option.WithNIOSUsername(username),
		option.WithNIOSPassword(password),
		option.WithDebug(true),
	)
}
