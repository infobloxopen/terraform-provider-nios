package cloud_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAwsuserDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_cloud_awsuser.test"
	resourceName := "nios_cloud_awsuser.test"
	var v cloud.Awsuser

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAwsuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsuserDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					}, testAccCheckAwsuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions
func testAccCheckAwsuserResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "access_key_id", dataSourceName, "result.0.access_key_id"),
		resource.TestCheckResourceAttrPair(resourceName, "account_id", dataSourceName, "result.0.account_id"),
		resource.TestCheckResourceAttrPair(resourceName, "govcloud_enabled", dataSourceName, "result.0.govcloud_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "last_used", dataSourceName, "result.0.last_used"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nios_user_name", dataSourceName, "result.0.nios_user_name"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
	}
}

func testAccAwsuserDataSourceConfigFilters() string {
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")

	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test" {
  access_key_id      = "%s"
  account_id         = "%s"
  govcloud_enabled   = false
  name               = "%s"
  secret_access_key  = "S1JGWfwcZWEY+uduuujhSkfpyhxigL9A/uaJ96mY"
}

data "nios_cloud_awsuser" "test" {
  filters = {
    name = nios_cloud_awsuser.test.name
  }
  depends_on = [nios_cloud_awsuser.test]
}
`, accessKeyId, accountId, name)
}
