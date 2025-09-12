// Retrieve a specific SNMP user by filters
data "nios_security_snmp_user" "get_snmpuser_with_filter" {
  filters = {
    name = "snmpuser_example_2"
  }
}

// Retrieve a specific SNMP user using Extensible Attributes
data "nios_security_snmp_user" "get_snmpuser_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SNMP users
data "nios_security_snmp_user" "get_all_snmpusers" {}
