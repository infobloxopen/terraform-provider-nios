// Retrieve a specific Syslog Endpoint by filters
data "nios_misc_syslog_endpoint" "get_syslog_endpoint_using_filters" {
  filters = {
    name = "syslogendpoint1"
  }
}
// Retrieve specific Syslog Endpoint using Extensible Attributes
data "nios_misc_syslog_endpoint" "get_syslog_endpoint_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Syslog Endpoints
data "nios_misc_syslog_endpoint" "get_all_syslog_endpoints" {}
