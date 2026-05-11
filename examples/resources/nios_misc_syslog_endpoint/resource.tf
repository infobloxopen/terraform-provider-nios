// Create Syslog Endpoint with Basic Fields
resource "nios_misc_syslog_endpoint" "syslogendpoint_1" {
  name                 = "syslogendpoint1"
  outbound_member_type = "GM"
  syslog_servers = [
    {
      address         = "10.1.1.1"
      port            = 514
      connection_type = "udp"
      format          = "formatted"
    }
  ]
}

// Create Syslog Endpoint with Additional Fields
resource "nios_misc_syslog_endpoint" "syslogendpoint_2" {
  name                 = "syslogendpoint2"
  outbound_member_type = "GM"
  syslog_servers = [
    {
      address         = "10.1.1.2"
      port            = 514
      connection_type = "udp"
      format          = "formatted"
    }
  ]
  wapi_user_name     = "admin"
  wapi_user_password = "Example-Admin123!"
  extattrs = {
    Site = "location-1"
  }
}
