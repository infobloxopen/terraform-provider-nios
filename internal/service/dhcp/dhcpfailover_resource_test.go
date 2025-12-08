package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
TODO: grid setup to run test cases
Grid Members - infoblox.localdomain, infoblox.member2 and infoblox.member1
*/
var readableAttributesForDhcpfailover = "association_type,comment,extattrs,failover_port,load_balance_split,max_client_lead_time,max_load_balance_delay,max_response_delay,max_unacked_updates,ms_association_mode,ms_enable_authentication,ms_enable_switchover_interval,ms_failover_mode,ms_failover_partner,ms_hotstandby_partner_role,ms_is_conflict,ms_previous_state,ms_server,ms_state,ms_switchover_interval,name,primary,primary_server_type,primary_state,recycle_leases,secondary,secondary_server_type,secondary_state,use_failover_port,use_ms_switchover_interval,use_recycle_leases"

func TestAccDhcpfailoverResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test"
	var v dhcp.Dhcpfailover

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverBasicConfig("dhcp_failover_test", "infoblox.localdomain", "infoblox.member1", "GRID", "GRID"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "dhcp_failover_test"),
					resource.TestCheckResourceAttr(resourceName, "primary", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "secondary", "infoblox.member1"),
					resource.TestCheckResourceAttr(resourceName, "primary_server_type", "GRID"),
					resource.TestCheckResourceAttr(resourceName, "secondary_server_type", "GRID"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "association_type", "GRID"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_failover_port", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ms_switchover_interval", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_failover_mode", "LOADBALANCE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_dhcpfailover.test"
	var v dhcp.Dhcpfailover

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpfailoverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpfailoverBasicConfig("dhcp_failover_test1", "infoblox.localdomain", "infoblox.member1", "GRID", "GRID"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					testAccCheckDhcpfailoverDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDhcpfailoverResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_comment"
	var v dhcp.Dhcpfailover
	dhcpfailoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverComment(dhcpfailoverName, primary, secondary, primaryServerType, secondaryServerType, "This is a test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverComment(dhcpfailoverName, primary, secondary, primaryServerType, secondaryServerType, "This is an updated test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated test comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_extattrs"
	var v dhcp.Dhcpfailover
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	dhcpfailoverName := acctest.RandomNameWithPrefix("dhcp_failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverExtAttrs(dhcpfailoverName, primary, secondary, primaryServerType, secondaryServerType, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverExtAttrs(dhcpfailoverName, primary, secondary, primaryServerType, secondaryServerType, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_FailoverPort(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_failover_port"
	var v dhcp.Dhcpfailover
	failoverPort := "647"
	updateFailoverPort := "648"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useFailoverPort := true

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverFailoverPort(failoverName, primary, secondary, primaryServerType, secondaryServerType, failoverPort, useFailoverPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_port", "647"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverFailoverPort(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateFailoverPort, useFailoverPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_port", "648"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_LoadBalanceSplit(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_load_balance_split"
	var v dhcp.Dhcpfailover
	loadBalanceSplit := "120"
	updateLoadBalanceSplit := "121"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverLoadBalanceSplit(failoverName, primary, secondary, primaryServerType, secondaryServerType, loadBalanceSplit),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "load_balance_split", "120"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverLoadBalanceSplit(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateLoadBalanceSplit),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "load_balance_split", "121"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MaxClientLeadTime(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_max_client_lead_time"
	var v dhcp.Dhcpfailover
	maxClientLeadTime := "4000"
	updateMaxClientLeadTime := "4001"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMaxClientLeadTime(failoverName, primary, secondary, primaryServerType, secondaryServerType, maxClientLeadTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_client_lead_time", "4000"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMaxClientLeadTime(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMaxClientLeadTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_client_lead_time", "4001"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MaxLoadBalanceDelay(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_max_load_balance_delay"
	var v dhcp.Dhcpfailover
	maxLoadBalanceDelay := "5000"
	updateMaxLoadBalanceDelay := "5001"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMaxLoadBalanceDelay(failoverName, primary, secondary, primaryServerType, secondaryServerType, maxLoadBalanceDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_load_balance_delay", "5000"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMaxLoadBalanceDelay(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMaxLoadBalanceDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_load_balance_delay", "5001"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MaxResponseDelay(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_max_response_delay"
	var v dhcp.Dhcpfailover
	maxResponseDelay := "6000"
	updateMaxResponseDelay := "6001"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMaxResponseDelay(failoverName, primary, secondary, primaryServerType, secondaryServerType, maxResponseDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_response_delay", "6000"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMaxResponseDelay(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMaxResponseDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_response_delay", "6001"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MaxUnackedUpdates(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_max_unacked_updates"
	var v dhcp.Dhcpfailover
	maxUnackedUpdates := "7000"
	updateMaxUnackedUpdates := "7001"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMaxUnackedUpdates(failoverName, primary, secondary, primaryServerType, secondaryServerType, maxUnackedUpdates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_unacked_updates", "7000"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMaxUnackedUpdates(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMaxUnackedUpdates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_unacked_updates", "7001"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MsEnableAuthentication(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_ms_enable_authentication"
	var v dhcp.Dhcpfailover
	msEnableAuthentication := "true"
	updateMsEnableAuthentication := "false"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMsEnableAuthentication(failoverName, primary, secondary, primaryServerType, secondaryServerType, msEnableAuthentication),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_enable_authentication", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMsEnableAuthentication(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMsEnableAuthentication),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_enable_authentication", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MsEnableSwitchoverInterval(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_ms_enable_switchover_interval"
	var v dhcp.Dhcpfailover
	msEnableSwitchoverInterval := "true"
	updateMsEnableSwitchoverInterval := "false"
	useMSSwitchoverIntervalUpdate := false
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useMSSwitchoverInterval := true

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMsEnableSwitchoverInterval(failoverName, primary, secondary, primaryServerType, secondaryServerType, msEnableSwitchoverInterval, useMSSwitchoverInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_enable_switchover_interval", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMsEnableSwitchoverInterval(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateMsEnableSwitchoverInterval, useMSSwitchoverIntervalUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_enable_switchover_interval", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MsFailoverMode(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_ms_failover_mode"
	var v dhcp.Dhcpfailover
	msFailoverMode := "HOTSTANDBY"
	updateMsFailoverMode := "LOADBALANCE"
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	name := acctest.RandomNameWithPrefix("failover")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMsFailoverMode(msFailoverMode, name, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_failover_mode", msFailoverMode),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMsFailoverMode(updateMsFailoverMode, name, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_failover_mode", updateMsFailoverMode),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MsHotstandbyPartnerRole(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_ms_hotstandby_partner_role"
	var v dhcp.Dhcpfailover
	name := acctest.RandomNameWithPrefix("failover")
	msHotstandbyPartnerRole := "ACTIVE"
	msFailoverMode := "HOTSTANDBY"
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMsHotstandbyPartnerRole(msHotstandbyPartnerRole, msFailoverMode, name, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_hotstandby_partner_role", "ACTIVE"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMsHotstandbyPartnerRole("PASSIVE", msFailoverMode, name, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_hotstandby_partner_role", "PASSIVE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_MsSwitchoverInterval(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_ms_switchover_interval"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	switchOverInterval := "300"
	changeSwitchOverInterval := "400"
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useMsSwitchoverInterval := true
	msEnableSwitchoverInterval := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverMsSwitchoverInterval(switchOverInterval, failoverName, primary, secondary, primaryServerType, secondaryServerType, useMsSwitchoverInterval, msEnableSwitchoverInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_switchover_interval", "300"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverMsSwitchoverInterval(changeSwitchOverInterval, failoverName, primary, secondary, primaryServerType, secondaryServerType, useMsSwitchoverInterval, msEnableSwitchoverInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_switchover_interval", "400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_name"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	updateFailoverName := acctest.RandomNameWithPrefix("failover_updated")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverName(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", failoverName),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverName(updateFailoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updateFailoverName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_Primary(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_primary"
	var v dhcp.Dhcpfailover
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	failoverName := acctest.RandomNameWithPrefix("failover")
	updatedPrimary := "infoblox.member2"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverPrimary(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "primary", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverPrimary(failoverName, updatedPrimary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "primary", "infoblox.member2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_PrimaryServerType(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_primary_server_type"
	var v dhcp.Dhcpfailover
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	secondaryServerType := "GRID"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primaryServerType := "GRID"
	updatedPrimaryServerType := "EXTERNAL"
	primaryServer := "172.28.82.25"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverPrimaryServerType(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "primary_server_type", "GRID"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverPrimaryServerType(failoverName, primaryServer, secondary, updatedPrimaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "primary_server_type", "EXTERNAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_recycle_leases"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	recycleLeases := true
	updateRecycleLeases := false
	useRecycleLeases := true
	useUpdateRecycleLeases := false

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverRecycleLeases(failoverName, primary, secondary, primaryServerType, secondaryServerType, recycleLeases, useRecycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverRecycleLeases(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateRecycleLeases, useUpdateRecycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_Secondary(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_secondary"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	secondaryUpdate := "infoblox.member2"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverSecondary(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary", secondary),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverSecondary(failoverName, primary, secondaryUpdate, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary", secondaryUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_SecondaryServerType(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_secondary_server_type"
	var v dhcp.Dhcpfailover
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	failoverName := acctest.RandomNameWithPrefix("failover")
	secondaryServer := "10.197.36.31"
	secondaryServerTypeUpdate := "EXTERNAL"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverSecondaryServerType(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_server_type", "GRID"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverSecondaryServerType(failoverName, primary, secondaryServer, primaryServerType, secondaryServerTypeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_server_type", "EXTERNAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_UseFailoverPort(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_use_failover_port"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useFailoverPort := "true"
	updateUseFailoverPort := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverUseFailoverPort(failoverName, primary, secondary, primaryServerType, secondaryServerType, useFailoverPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_failover_port", useFailoverPort),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverUseFailoverPort(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateUseFailoverPort),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_failover_port", updateUseFailoverPort),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_UseMsSwitchoverInterval(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_use_ms_switchover_interval"
	var v dhcp.Dhcpfailover
	useMsSwitchoverInterval := "true"
	updateUseMsSwitchoverInterval := "false"
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useSwitchoverInterval := "true"
	useUpdateSwitchoverInterval := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverUseMsSwitchoverInterval(failoverName, primary, secondary, primaryServerType, secondaryServerType, useMsSwitchoverInterval, useSwitchoverInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_switchover_interval", useMsSwitchoverInterval),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverUseMsSwitchoverInterval(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateUseMsSwitchoverInterval, useUpdateSwitchoverInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_switchover_interval", updateUseMsSwitchoverInterval),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpfailoverResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpfailover.test_use_recycle_leases"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	useRecycleLeases := "true"
	updateUseRecycleLeases := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpfailoverUseRecycleLeases(failoverName, primary, secondary, primaryServerType, secondaryServerType, useRecycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", useRecycleLeases),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpfailoverUseRecycleLeases(failoverName, primary, secondary, primaryServerType, secondaryServerType, updateUseRecycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", updateUseRecycleLeases),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDhcpfailoverExists(ctx context.Context, resourceName string, v *dhcp.Dhcpfailover) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			DhcpfailoverAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDhcpfailover).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDhcpfailoverResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDhcpfailoverResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDhcpfailoverDestroy(ctx context.Context, v *dhcp.Dhcpfailover) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			DhcpfailoverAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDhcpfailover).
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

func testAccCheckDhcpfailoverDisappears(ctx context.Context, v *dhcp.Dhcpfailover) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			DhcpfailoverAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDhcpfailoverBasicConfig(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test" {
    name = %q
    primary = %q
    secondary = %q
    primary_server_type = %q
    secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverComment(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_comment" {
    comment = %q
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, comment, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverExtAttrs(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"

	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_extattrs" {
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    extattrs = %s
}
`, name, primary, secondary, primaryServerType, secondaryServerType, extattrsStr)
}

func testAccDhcpfailoverFailoverPort(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, failoverPort string, useFailoverPort bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_failover_port" {
    failover_port = %q
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
	use_failover_port = %t
}
`, failoverPort, name, primary, secondary, primaryServerType, secondaryServerType, useFailoverPort)
}

func testAccDhcpfailoverLoadBalanceSplit(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, loadBalanceSplit string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_load_balance_split" {
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    load_balance_split = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, loadBalanceSplit)
}

func testAccDhcpfailoverMaxClientLeadTime(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, maxClientLeadTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_max_client_lead_time" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    max_client_lead_time = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, maxClientLeadTime)
}

func testAccDhcpfailoverMaxLoadBalanceDelay(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, maxLoadBalanceDelay string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_max_load_balance_delay" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    max_load_balance_delay = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, maxLoadBalanceDelay)
}

func testAccDhcpfailoverMaxResponseDelay(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, maxResponseDelay string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_max_response_delay" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    max_response_delay = %q
}	
`, name, primary, secondary, primaryServerType, secondaryServerType, maxResponseDelay)
}

func testAccDhcpfailoverMaxUnackedUpdates(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, maxUnackedUpdates string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_max_unacked_updates" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    max_unacked_updates = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, maxUnackedUpdates)
}

func testAccDhcpfailoverMsEnableAuthentication(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, msEnableAuthentication string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_ms_enable_authentication" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    ms_enable_authentication = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, msEnableAuthentication)
}

func testAccDhcpfailoverMsEnableSwitchoverInterval(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, msEnableSwitchoverInterval string, useMSSwitchoverInterval bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_ms_enable_switchover_interval" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    ms_enable_switchover_interval = %q
	use_ms_switchover_interval = %t
}
`, name, primary, secondary, primaryServerType, secondaryServerType, msEnableSwitchoverInterval, useMSSwitchoverInterval)
}

func testAccDhcpfailoverMsFailoverMode(msFailoverMode string, name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_ms_failover_mode" {
    ms_failover_mode = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
	name = %q
}
`, msFailoverMode, primary, secondary, primaryServerType, secondaryServerType, name)
}

func testAccDhcpfailoverMsHotstandbyPartnerRole(msHotstandbyPartnerRole string, msFailoverMode string, name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_ms_hotstandby_partner_role" {
    ms_hotstandby_partner_role = %q
	ms_failover_mode = %q
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, msHotstandbyPartnerRole, msFailoverMode, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverMsSwitchoverInterval(msSwitchoverInterval string, name string, primary string, secondary string, primaryServerType string, secondaryServerType string, useMsSwitchoverInterval bool, msEnableSwitchoverInterval string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_ms_switchover_interval" {
    ms_switchover_interval = %q
	name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
	use_ms_switchover_interval = %t
	ms_enable_switchover_interval = %q
}
`, msSwitchoverInterval, name, primary, secondary, primaryServerType, secondaryServerType, useMsSwitchoverInterval, msEnableSwitchoverInterval)
}

func testAccDhcpfailoverName(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_name" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverPrimary(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_primary" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverPrimaryServerType(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_primary_server_type" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverRecycleLeases(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, recycleLeases bool, useRecycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_recycle_leases" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    recycle_leases = %t
	use_recycle_leases = %t
}
`, name, primary, secondary, primaryServerType, secondaryServerType, recycleLeases, useRecycleLeases)
}

func testAccDhcpfailoverSecondary(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_secondary" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverSecondaryServerType(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_secondary_server_type" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverUseFailoverPort(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, useFailoverPort string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_use_failover_port" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    use_failover_port = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, useFailoverPort)
}

func testAccDhcpfailoverUseMsSwitchoverInterval(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, useMsSwitchoverInterval string, useEnableSwitchoverInterval string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_use_ms_switchover_interval" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    use_ms_switchover_interval = %q
	ms_enable_switchover_interval = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, useMsSwitchoverInterval, useEnableSwitchoverInterval)
}

func testAccDhcpfailoverUseRecycleLeases(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test_use_recycle_leases" {
    name = %q
	primary = %q
	secondary = %q
	primary_server_type = %q
	secondary_server_type = %q
    use_recycle_leases = %q
}
`, name, primary, secondary, primaryServerType, secondaryServerType, useRecycleLeases)
}
