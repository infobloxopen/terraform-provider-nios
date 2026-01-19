// Create an SNMP User with Basic Fields
resource "nios_security_snmp_user" "snmpuser_basic_fields" {
  name                    = "snmpuser_example_1"
  authentication_protocol = "NONE"
  privacy_protocol        = "NONE"
}

// Create an SNMP User with Additional Fields
resource "nios_security_snmp_user" "snmpuser_additional_fields" {
  name                    = "snmpuser_example_2"
  authentication_protocol = "SHA"
  authentication_password = "AuthPassword@123"
  privacy_protocol        = "DES"
  privacy_password        = "PrivacyPassword@123"
  comment                 = "Example SNMP User"
  extattrs = {
    Site = "location-1"
  }
}
