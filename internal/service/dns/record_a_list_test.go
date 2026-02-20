package dns_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordAList_basic(t *testing.T) {
	//var resourceName = "nios_dns_record_a.test"
	//var v dns.RecordA
	//name := acctest.RandomName() + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: basicConfig(),
			},
			{
				Query:  true,
				Config: testAccRecordAListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_a.test", 3),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func basicConfig() string {
	return `
	terraform {
  required_providers {
    nios = {
      source  = "registry.terraform.io/infobloxopen/nios"
      version = "1.0.0"
    }
  }
  required_version = ">= 1.8.0"
}
`
}

func testAccRecordAListBasicConfig() string {
	return `
list "nios_dns_record_a" "test" {
  provider = nios
  include_resource = true
}
`
}
