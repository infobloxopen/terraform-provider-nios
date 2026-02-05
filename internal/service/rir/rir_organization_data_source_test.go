
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationDataSourceConfigFilters(),
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
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationDataSourceConfigExtAttrFilters( acctest.RandomName()),
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

func testAccCheckRirOrganizationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "result.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "maintainer", dataSourceName, "result.0.maintainer"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "password", dataSourceName, "result.0.password"),
        resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
        resource.TestCheckResourceAttrPair(resourceName, "sender_email", dataSourceName, "result.0.sender_email"),
    }
}

func testAccRirOrganizationDataSourceConfigFilters() string {
	return `
resource "nios_rir_organization" "test" {
}

data "nios_rir_organization" "test" {
  filters = {
	 = nios_rir_organization.test.
  }
}
`
}

func testAccRirOrganizationDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_rir_organization" "test" {
  extattrfilters = {
	Site = nios_rir_organization.test.extattrs.Site
  }
}
`,extAttrsValue)
}

