package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNetworkview = "associated_dns_views,associated_members,cloud_info,comment,ddns_dns_view,ddns_zone_primaries,extattrs,internal_forward_zones,is_default,mgm_private,ms_ad_user_data,name,remote_forward_zones,remote_reverse_zones"

// TODO: Prerequisites
// Pre-provision CP Member with Cloud API license
// CP Member Name : "infoblox.cloudmem"
// CP Member IPv4 Address : "172.172.172.172"
// CP Member IPv6 : "2001::123"
// Create Auth Zone: "first.com", "second.com" with Grid Primary: "infoblox.cloudmem"

func TestAccNetworkviewResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_networkview.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkviewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					testAccCheckNetworkviewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkviewResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_cloud_info"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")
	member_name := "infoblox.cloudmem"
	member_ipv4 := "172.172.172.172"
	member_ipv6 := "2001::123"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewCloudInfo(name, member_name, member_ipv4, member_ipv6),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "ROOT"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_member.name", member_name),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_member.ipv4addr", member_ipv4),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_member.ipv6addr", member_ipv6),
				),
			},
			// Update and Read
			// TODO: unset the delegated_member using 'null' and delegate to a different member
			// {
			// 	Config: testAccNetworkviewCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
			// 	),
			// },
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_comment"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewComment(name, "This is a new network view"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new network view"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewComment(name, "This is an modified network view"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an modified network view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_DdnsDnsView(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ddns_dns_view"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")
	DdnsDnsView := "default." + name

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewDdnsDnsView(name, "null"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_dns_view", DdnsDnsView),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewDdnsDnsView(name, "\""+DdnsDnsView+"\""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_dns_view", DdnsDnsView),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_DdnsZonePrimaries(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ddns_zone_primaries"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")
	DdnsZonePrimariesCreate := `[
        {
        dns_grid_primary ="infoblox.cloudmem",
        zone_match = "GRID",
        dns_grid_zone = {
            ref= "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default"
            }
        }
	]`

	DdnsZonePrimariesUpdate := `[
        {
        dns_grid_primary ="infoblox.cloudmem",
        zone_match = "GRID",
        dns_grid_zone = {
            ref= "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default"
            }
        },
        {
        dns_grid_primary ="infoblox.cloudmem",
        zone_match = "GRID",
        dns_grid_zone = {
            ref= "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5zZWNvbmQ:second.com/default"
            }
        },
	]`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewDdnsZonePrimaries(name, DdnsZonePrimariesCreate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.0.dns_grid_primary", "infoblox.cloudmem"),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.0.zone_match", "GRID"),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.0.dns_grid_zone.ref", "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewDdnsZonePrimaries(name, DdnsZonePrimariesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.1.dns_grid_primary", "infoblox.cloudmem"),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.1.zone_match", "GRID"),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries.1.dns_grid_zone.ref", "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5zZWNvbmQ:second.com/default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_extattrs"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("nview")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewExtAttrs(name, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewExtAttrs(name, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO: Requires UDDI configs
// func TestAccNetworkviewResource_FederatedRealms(t *testing.T) {
// 	var resourceName = "nios_ipam_networkview.test_federated_realms"
// 	var v ipam.Networkview
// 	name := acctest.RandomNameWithPrefix("test-nview")
// 	FederatedRealmsCreate := `[
//         {
//         id = "11",
//         name = "federated_realm_1",
// 		},
// 	]`
// 	FederatedRealmsUpdate := `[
// 		{
// 		id = "22",
// 		name = "federated_realm_2",
// 		},
// 	]`

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkviewFederatedRealms(name, FederatedRealmsCreate),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms.0.id", "11"),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms.0.name", "federated_realm_1"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkviewFederatedRealms(name, FederatedRealmsUpdate),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms.0.id", "22"),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms.0.name", "federated_realm_2"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// TODO: Prerequisite zones to be created after network view creation
// func TestAccNetworkviewResource_InternalForwardZones(t *testing.T) {
// 	var resourceName = "nios_ipam_networkview.test_internal_forward_zones"
// 	var v ipam.Networkview
// 	name := acctest.RandomNameWithPrefix("test-nview")
// 	InternalForwardZonesCreate := []string{"\"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default\""}
// 	InternalForwardZonesUpdate := []string{"\"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default\"",
// 		"\"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5zZWNvbmQ:second.com/default\""}

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkviewInternalForwardZonesBasic(name),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkviewInternalForwardZones(name, InternalForwardZonesCreate),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "internal_forward_zones.0", InternalForwardZonesCreate[0]),
// 				),
// 			},
// 			{
// 				Config: testAccNetworkviewInternalForwardZones(name, InternalForwardZonesUpdate),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "internal_forward_zones.0", InternalForwardZonesUpdate[0]),
// 					resource.TestCheckResourceAttr(resourceName, "internal_forward_zones.1", InternalForwardZonesUpdate[1]),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func TestAccNetworkviewResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_mgm_private"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")
	MgmPrivateCreate := false
	MgmPrivateUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewMgmPrivate(name, MgmPrivateCreate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewMgmPrivate(name, MgmPrivateUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_name"
	var v ipam.Networkview
	Name := acctest.RandomNameWithPrefix("test-nview")
	NewName := acctest.RandomNameWithPrefix("test-nview")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewName(Name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", Name),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewName(NewName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", NewName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_RemoteForwardZones(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_remote_forward_zones"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")
	RemoteForwardZonesCreate := `[
	{
		fqdn           = "fwdzone1.com"
		key_type       = "NONE"
		server_address = "192.168.12.12"
	}
	]`

	RemoteForwardZonesUpdate := `[
	{
		fqdn           = "fwdzone2.com"
		key_type       = "NONE"
		server_address = "192.168.12.13"
	},
	{
		fqdn           = "fwdzone3.com"
		key_type       = "TSIG"
		server_address = "192.168.12.14"
		tsig_key_name  = "tsigkey"
		tsig_key_alg   = "HMAC-SHA256"
		tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
	}
	]`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewRemoteForwardZones(name, RemoteForwardZonesCreate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.fqdn", "fwdzone1.com"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.key_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.server_address", "192.168.12.12"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewRemoteForwardZones(name, RemoteForwardZonesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.fqdn", "fwdzone2.com"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.key_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.0.server_address", "192.168.12.13"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.fqdn", "fwdzone3.com"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.key_type", "TSIG"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.server_address", "192.168.12.14"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.tsig_key_name", "tsigkey"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.tsig_key_alg", "HMAC-SHA256"),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones.1.tsig_key", "dGhpc2lzdGVzdHRzaWdrZXk="),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_RemoteReverseZones(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_remote_reverse_zones"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-nview")

	RemoteReverseZonesCreate := `[
	{
		fqdn           = "0.168.192.in-addr.arpa"
		key_type       = "NONE"
		server_address = "192.168.12.12"
	}
	]`

	RemoteReverseZonesUpdate := `[
	{
		fqdn           = "2.168.192.in-addr.arpa"
		key_type       = "NONE"
		server_address = "192.168.12.13"
	},
	{
		fqdn           = "1.168.192.in-addr.arpa"
		key_type       = "TSIG"
		server_address = "192.168.12.14"
		tsig_key_name  = "tsigkey"
		tsig_key_alg   = "HMAC-SHA256"
		tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
	}
	]`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewRemoteReverseZones(name, RemoteReverseZonesCreate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.fqdn", "0.168.192.in-addr.arpa"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.key_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.server_address", "192.168.12.12"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewRemoteReverseZones(name, RemoteReverseZonesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.fqdn", "2.168.192.in-addr.arpa"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.key_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.0.server_address", "192.168.12.13"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.fqdn", "1.168.192.in-addr.arpa"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.key_type", "TSIG"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.server_address", "192.168.12.14"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.tsig_key_name", "tsigkey"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.tsig_key_alg", "HMAC-SHA256"),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones.1.tsig_key", "dGhpc2lzdGVzdHRzaWdrZXk="),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworkviewExists(ctx context.Context, resourceName string, v *ipam.Networkview) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetworkview).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworkviewResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworkviewResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworkviewDestroy(ctx context.Context, v *ipam.Networkview) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetworkview).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckNetworkviewDisappears(ctx context.Context, v *ipam.Networkview) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkviewBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test" {
	name = %q
}
`, name)
}

func testAccNetworkviewRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNetworkviewCloudInfo(name, member_name, member_ipv4, member_ipv6 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_cloud_info" {
	name = %q
    cloud_info = {
		delegated_member = {
			name 	 = %q
			ipv4addr = %q
			ipv6addr = %q
		}
	}
}
`, name, member_name, member_ipv4, member_ipv6)
}

func testAccNetworkviewComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_comment" {
	name 	= %q
    comment = %q
}
`, name, comment)
}

func testAccNetworkviewDdnsDnsView(name, ddnsDnsView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ddns_dns_view" {
	name 		  = %q
    ddns_dns_view = %s
}
`, name, ddnsDnsView)
}

func testAccNetworkviewDdnsZonePrimaries(name, ddnsZonePrimaries string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ddns_zone_primaries" {
	name = %q
    ddns_zone_primaries = %s
}
`, name, ddnsZonePrimaries)
}

func testAccNetworkviewExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccNetworkviewFederatedRealms(name, federatedRealms string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_federated_realms" {
	name = %q
    federated_realms = %s
}
`, name, federatedRealms)
}

func testAccNetworkviewInternalForwardZonesBasic(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_internal_forward_zones" {
    name = %q
}
`, name)
}

func testAccNetworkviewInternalForwardZones(name string, internalForwardZones []string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_internal_forward_zones" {
	name = %q
    internal_forward_zones = %s
}
`, name, internalForwardZones)
}

func testAccNetworkviewMgmPrivate(name string, mgmPrivate bool) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_mgm_private" {
	name = %q
    mgm_private = %t
}
`, name, mgmPrivate)
}

func testAccNetworkviewMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccNetworkviewName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_name" {
    name = %q
}
`, name)
}

func testAccNetworkviewRemoteForwardZones(name, remoteForwardZones string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_remote_forward_zones" {
	name 				 = %q
    remote_forward_zones = %s
}
`, name, remoteForwardZones)
}

func testAccNetworkviewRemoteReverseZones(name, remoteReverseZones string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_remote_reverse_zones" {
	name 				 = %q
    remote_reverse_zones = %s
}
`, name, remoteReverseZones)
}
