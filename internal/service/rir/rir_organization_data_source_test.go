package rir_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/rir"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRirOrganizationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rir_organization.test"
	resourceName := "nios_rir_organization.test"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationDataSourceConfigFilters(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					}, testAccCheckRirOrganizationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRirOrganizationDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rir_organization.test"
	resourceName := "nios_rir_organization.test"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationDataSourceConfigExtAttrFilters(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					}, testAccCheckRirOrganizationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRirOrganizationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "result.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "maintainer", dataSourceName, "result.0.maintainer"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "sender_email", dataSourceName, "result.0.sender_email"),
	}
}

func testAccRirOrganizationDataSourceConfigFilters(extAttrs map[string]string, id, maintainer, name, password, rir, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
  extattrs = %s
  id = %q
  maintainer = %q
  name = %q
  password = %q
  rir = %q
  sender_email = %q
}

data "nios_rir_organization" "test" {
  filters = {
	name = nios_rir_organization.test.name
  }
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationDataSourceConfigExtAttrFilters(extAttrs map[string]string, id, maintainer, name, password, rir, senderEmail, extAttrsValue string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
  extattrs = %s
  id = %q
  maintainer = %q
  name = %q
  password = %q
  rir = %q
  sender_email = %q
}

data "nios_rir_organization" "test" {
  extattrfilters = {
    "RIPE Email" = nios_rir_organization.test.extattrs["RIPE Email"]
  }
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}
