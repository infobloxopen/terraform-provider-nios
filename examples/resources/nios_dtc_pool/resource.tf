//create a DTC pool with basic parameters 
resource "nios_dtc_pool" "dtc_pool1"{
    name = "dtc_pool"
    lb_preferred_method="ROUND_ROBIN"
}

//create a DTC pool with maximal parameters 
resource "nios_dtc_pool" "dtc_pool2"{
    name  = "dtc_pool2"
    lb_preferred_method = "TOPOLOGY"
    lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
    comment = "DTC pool creation"
    extattrs = {
    Site = "Siteblr"
  }
    servers = [
        {
            server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"
            ratio = 34
        },
        {
            server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com"
            ratio = 55
        }
    ]
    monitors = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http","dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
    lb_alternate_method = "DYNAMIC_RATIO"
    lb_dynamic_ratio_alternate = {
        method = "ROUND_TRIP_DELAY"
        monitor_weighing = "RATIO"
        invert_monitor_metric = true
        monitor = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"
        monitor_metric = ".2"
    }
    auto_consolidated_monitors = true
    disable = false 
}

//Consolidated monitors can only be added if the pool is associated with an LBDN.
resource "nios_dtc_pool" "dtc_pool3"{
    name = "dtc_pool3"
    lb_preferred_method = "ROUND_ROBIN"
    monitors = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http","dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
    consolidated_monitors =[
    {
      monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      members = ["infoblox.localdomain"]
      availability = "ALL"
      full_health_communication = true
    }
  ]
}