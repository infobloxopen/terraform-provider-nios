package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMacfilteraddress = "authentication_time,comment,expiration_time,extattrs,filter,fingerprint,guest_custom_field1,guest_custom_field2,guest_custom_field3,guest_custom_field4,guest_email,guest_first_name,guest_last_name,guest_middle_name,guest_phone,is_registered_user,mac,never_expires,reserved_for_infoblox,username"

func TestAccMacfilteraddressResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test"
	var v dhcp.Macfilteraddress
	mac := "00:1A:2B:3C:3D:5E"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressBasicConfig(filter, mac),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter", filter),
					resource.TestCheckResourceAttr(resourceName, "mac", "00:1A:2B:3C:3D:5E"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "never_expires", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_registered_user", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_email", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_first_name", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_last_name", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_middle_name", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_phone", ""),
					resource.TestCheckResourceAttr(resourceName, "reserved_for_infoblox", ""),
					resource.TestCheckResourceAttr(resourceName, "username", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field1", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field2", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field3", ""),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field4", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_macfilteraddress.test"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMacfilteraddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMacfilteraddressBasicConfig(filter, mac),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					testAccCheckMacfilteraddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMacfilteraddressResource_AuthenticationTime(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_authentication_time"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressAuthenticationTime(filter, mac, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_time", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressAuthenticationTime(filter, mac, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_time", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_comment"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressComment(filter, mac, "Filter Mac Address Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Filter Mac Address Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressComment(filter, mac, "Filter Mac Address Comment Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Filter Mac Address Comment Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_ExpirationTime(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_expiration_time"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressExpirationTime(filter, mac, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration_time", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressExpirationTime(filter, mac, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration_time", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_extattrs"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	extAttrs1 := acctest.RandomName()
	extAttrs2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressExtAttrs(filter, mac, map[string]any{"Site": extAttrs1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs1),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressExtAttrs(filter, mac, map[string]any{"Site": extAttrs2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_Filter(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_filter"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter1 := acctest.RandomNameWithPrefix("mac-filter")
	filter2 := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressFilter(filter1, filter2, mac, "parent_filter_mac1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter", filter1),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressFilter(filter1, filter2, mac, "parent_filter_mac2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter", filter2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestCustomField1(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_custom_field1"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestCustomField1(filter, mac, "guest-user-field1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field1", "guest-user-field1"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestCustomField1(filter, mac, "guest-user-field1-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field1", "guest-user-field1-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestCustomField2(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_custom_field2"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestCustomField2(filter, mac, "guest-user-field2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field2", "guest-user-field2"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestCustomField2(filter, mac, "guest-user-field2-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field2", "guest-user-field2-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestCustomField3(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_custom_field3"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestCustomField3(filter, mac, "guest-user-fields3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field3", "guest-user-fields3"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestCustomField3(filter, mac, "guest-user-field3-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field3", "guest-user-field3-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestCustomField4(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_custom_field4"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestCustomField4(filter, mac, "guest-user-field4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field4", "guest-user-field4"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestCustomField4(filter, mac, "guest-user-field4-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_custom_field4", "guest-user-field4-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestEmail(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_email"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestEmail(filter, mac, "abc@xyz.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_email", "abc@xyz.com"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestEmail(filter, mac, "aaa@xyz.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_email", "aaa@xyz.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestFirstName(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_first_name"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestFirstName(filter, mac, "firstname"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_first_name", "firstname"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestFirstName(filter, mac, "firstname-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_first_name", "firstname-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestLastName(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_last_name"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestLastName(filter, mac, "lastname"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_last_name", "lastname"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestLastName(filter, mac, "lastname-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_last_name", "lastname-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestMiddleName(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_middle_name"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestMiddleName(filter, mac, "middlename"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_middle_name", "middlename"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestMiddleName(filter, mac, "middlename-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_middle_name", "middlename-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_GuestPhone(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_guest_phone"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressGuestPhone(filter, mac, "1234567890"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_phone", "1234567890"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressGuestPhone(filter, mac, "0987654321"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guest_phone", "0987654321"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_Mac(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_mac"
	var v dhcp.Macfilteraddress
	mac1 := "00:1a:2b:3c:3d:5e"
	mac2 := "00:1a:2b:3c:3d:6f"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressMac(filter, mac1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", mac1),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressMac(filter, mac2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", mac2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_NeverExpires(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_never_expires"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressNeverExpires(filter, mac, true, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "never_expires", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressNeverExpires(filter, mac, false, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "never_expires", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_ReservedForInfoblox(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_reserved_for_infoblox"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressReservedForInfoblox(filter, mac, "Reserved_For_Infoblox_Value"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_for_infoblox", "Reserved_For_Infoblox_Value"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressReservedForInfoblox(filter, mac, "Updated_Reserved_For_Infoblox_Value"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_for_infoblox", "Updated_Reserved_For_Infoblox_Value"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMacfilteraddressResource_Username(t *testing.T) {
	var resourceName = "nios_dhcp_macfilteraddress.test_username"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMacfilteraddressUsername(filter, mac, "user1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "user1"),
				),
			},
			// Update and Read
			{
				Config: testAccMacfilteraddressUsername(filter, mac, "user2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "user2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMacfilteraddressExists(ctx context.Context, resourceName string, v *dhcp.Macfilteraddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			MacfilteraddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMacfilteraddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMacfilteraddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMacfilteraddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMacfilteraddressDestroy(ctx context.Context, v *dhcp.Macfilteraddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			MacfilteraddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMacfilteraddress).
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

func testAccCheckMacfilteraddressDisappears(ctx context.Context, v *dhcp.Macfilteraddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			MacfilteraddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMacfilteraddressBasicConfig(filter, mac string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test" {
	mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
}
`, mac)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressAuthenticationTime(filter, mac string, authenticationTime int32) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_authentication_time" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    authentication_time = %d
}
`, mac, authenticationTime)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressComment(filter, mac, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_comment" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    comment = %q
}
`, mac, comment)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressExpirationTime(filter, mac string, expirationTime int32) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_expiration_time" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    expiration_time = %d
}
`, mac, expirationTime)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressExtAttrs(filter, mac string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_extattrs" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    extattrs = %s
}
`, mac, extAttrsStr)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressFilter(filter1, filter2, mac, parentFilterName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_filter" {
    mac    = %q
    filter = nios_dhcp_filtermac.%s.name
}
`, mac, parentFilterName)
	return strings.Join([]string{testAccBaseWithTwoMacFilters(filter1, filter2), config}, "")
}

func testAccMacfilteraddressGuestCustomField1(filter, mac, guestCustomField1 string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_custom_field1" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_custom_field1 = %q
}
`, mac, guestCustomField1)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestCustomField2(filter, mac, guestCustomField2 string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_custom_field2" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_custom_field2 = %q
}
`, mac, guestCustomField2)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestCustomField3(filter, mac, guestCustomField3 string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_custom_field3" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_custom_field3 = %q
}
`, mac, guestCustomField3)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestCustomField4(filter, mac, guestCustomField4 string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_custom_field4" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_custom_field4 = %q
}
`, mac, guestCustomField4)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestEmail(filter, mac, guestEmail string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_email" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_email = %q
}
`, mac, guestEmail)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestFirstName(filter, mac, guestFirstName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_first_name" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_first_name = %q
}
`, mac, guestFirstName)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestLastName(filter, mac, guestLastName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_last_name" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_last_name = %q
}
`, mac, guestLastName)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestMiddleName(filter, mac, guestMiddleName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_middle_name" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_middle_name = %q
}
`, mac, guestMiddleName)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressGuestPhone(filter, mac, guestPhone string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_guest_phone" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    guest_phone = %q
}
`, mac, guestPhone)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressMac(filter, mac string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_mac" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
}
`, mac)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressNeverExpires(filter, mac string, neverExpires bool, expirationTime int32) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_never_expires" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    never_expires = %t
    expiration_time = %d
}
`, mac, neverExpires, expirationTime)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressReservedForInfoblox(filter, mac, reservedForInfoblox string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_reserved_for_infoblox" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    reserved_for_infoblox = %q
}
`, mac, reservedForInfoblox)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressUsername(filter, mac, username string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test_username" {
    mac    = %q
    filter = nios_dhcp_filtermac.parent_filter_mac.name
    username = %q
}
`, mac, username)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccBaseWithMacFilter(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "parent_filter_mac" {
    name = %q
}
`, name)
}

func testAccBaseWithTwoMacFilters(filter1, filter2 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "parent_filter_mac1" {
    name = %q
}
resource "nios_dhcp_filtermac" "parent_filter_mac2" {
    name = %q
}
`, filter1, filter2)
}
