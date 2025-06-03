resource "nios_resource_nios_RecordA" "create_record" {
  name     = "example_test.test.com"
  ipv4addr = "10.20.1.2"
  view     = "default"
  extattrs = {
    Site = {
      value = "Siteblr"
    }
  }
}

resource "nios_resource_nios_DtcPool" "create_pool"{
  name = "dtc_pool__new1"
  lb_preferred_method = "ROUND_ROBIN"
  disable = true
  availability = "QUORUM"
  quorum = 1
  monitors =["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
  servers = [{
    server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVy:test-server"
    ratio = 3
  },
  {
    server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVy:test-server"
    ratio = 3
  }
  ]
  
}

resource "nios_resource_nios_DtcPool" "create_pool2"{
  name = "dtc_pool__new45"
  lb_preferred_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_preferred = {
    method = "ROUND_TRIP_DELAY"
    monitor_weighing = "RATIO"
    invert_monitor_metric= false 
    monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
    monitor_metric = ".0"
  }
  monitors =["dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp","dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
}

resource "nios_resource_nios_DtcPool" "create_pool3"{
  name = "dtc_pool__new46"
  lb_preferred_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_preferred = {
    method = "ROUND_TRIP_DELAY"
    monitor_weighing = "RATIO"
    invert_monitor_metric= false 
    monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
    monitor_metric = ".0"
  }
  servers = [
    {
    server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3QuY29t:test.com"
    ratio = 3
  },
    {
    server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com"
    ratio = 3
  },
  {
    server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"
    ratio = 3
  },
  ]
  monitors =["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http","dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
}

resource "nios_resource_nios_DtcPool" "create_pool4"{
  name = "dtc_pool__new47"
  lb_preferred_method = "ROUND_ROBIN"
  monitors =["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
  auto_consolidated_monitors = false
  consolidated_monitors =[
    {
      monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      members = ["infoblox.localdomain"]
      availability= "ALL"
    }
  ]
  disable = true 
  quorum = 4
}

resource "nios_resource_nios_DtcPool" "create_pool4"{
  name = "dtc_pool__new47"
  lb_preferred_method = "ROUND_ROBIN"
  monitors =["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http",""]
  auto_consolidated_monitors = true
  quorum = 4
}