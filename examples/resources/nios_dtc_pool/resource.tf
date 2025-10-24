// Create a DTC pool with basic fields 
resource "nios_dtc_pool" "dtc_pool1" {
  name                = "dtc_pool"
  lb_preferred_method = "ROUND_ROBIN"
}

// Create DTC Server (Required as Parent)
resource "nios_dtc_server" "dtc_server_1" {
  name = "example-server-for-pool1.com"
  host = "2.3.3.4"
}

resource "nios_dtc_server" "dtc_server_2" {
  name = "example-server-for-pool2.com"
  host = "2.3.3.4"
}

// Create a DTC pool with additional fields
resource "nios_dtc_pool" "dtc_pool2" {
  name                = "dtc_pool2"
  lb_preferred_method = "TOPOLOGY"
  // The topology ruleset used here must have any one of the server members configured in its topology members
  lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
  comment               = "DTC pool creation"
  extattrs = {
    Site = "location-1"
  }
  servers = [
    {
      server = nios_dtc_server.dtc_server_1.ref
      ratio  = 34
    },
    {
      server = nios_dtc_server.dtc_server_2.ref
      ratio  = 55
    }
  ]
  monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
  lb_alternate_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_alternate = {
    method                = "ROUND_TRIP_DELAY"
    monitor_weighing      = "RATIO"
    invert_monitor_metric = true
    //OID must (0.2) must be configured on the SNMP monitor used
    monitor        = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"
    monitor_metric = ".2"
  }
  auto_consolidated_monitors = true
  disable                    = false
  availability               = "QUORUM"
  quorum                     = 1
  ttl                        = 23
  use_ttl                    = true
}

// Create a DTC pool with consolidated monitors 
// Steps:
// Create a DTC pool without consolidated monitors 
// Associate the DTC pool with a DTC LBDN that has a zone with infoblox.localdomain member 
// Update this DTC pool to add consolidated monitors 
// Pre-requisite: A DTC LBDN with a zone that has infoblox.localdomain as member should exist and the pool should be associated with that LBDN
resource "nios_dtc_pool" "dtc_pool3" {
  name                = "dtc_pool3"
  lb_preferred_method = "ROUND_ROBIN"
  monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
  consolidated_monitors = [
    {
      monitor                   = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      members                   = ["infoblox.localdomain"]
      availability              = "ALL"
      full_health_communication = true
    }
  ]
}
