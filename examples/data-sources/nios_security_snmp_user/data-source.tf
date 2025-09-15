// Retrieve a specific SNMP user by filters
data "nios_security_snmp_user" "get_snmpuser_using_filter" {
  filters = {
    name = "snmpuser_example_2"
  }
}

// Retrieve specific SNMP users using Extensible Attributes
data "nios_security_snmp_user" "get_snmpuser_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SNMP users
data "nios_security_snmp_user" "get_all_snmpusers" {}
