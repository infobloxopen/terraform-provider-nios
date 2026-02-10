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
				Config: testAccRirOrganizationDataSourceConfigFilters("ID_REPLACE_ME", "MAINTAINER_REPLACE_ME", "NAME_REPLACE_ME", "PASSWORD_REPLACE_ME", "RIR_REPLACE_ME", "SENDER_EMAIL_REPLACE_ME"),
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
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "result.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "maintainer", dataSourceName, "result.0.maintainer"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "password", dataSourceName, "result.0.password"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "sender_email", dataSourceName, "result.0.sender_email"),
	}
}

func testAccRirOrganizationDataSourceConfigFilters(id, maintainer, name, password, rir, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
  id = %q
  maintainer = %q
  name = %q
  password = %q
  rir = %q
  sender_email = %q
}

data "nios_rir_organization" "test" {
  filters = {
	id = nios_rir_organization.test.id
  }
}
`, id, maintainer, name, password, rir, senderEmail)
}
