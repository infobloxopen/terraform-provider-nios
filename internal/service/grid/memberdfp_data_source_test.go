package grid_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMemberdfpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_memberdfp.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberdfpDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "result.#"),
					// Verify we got at least one result
					resource.TestCheckResourceAttrWith(dataSourceName, "result.#", func(value string) error {
						count, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						if count < 1 {
							return fmt.Errorf("expected at least 1 result, got %d", count)
						}
						return nil
					}),
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccMemberdfpDataSourceConfigFilters() string {
	return `
data "nios_grid_memberdfp" "test" {
}
`
}
