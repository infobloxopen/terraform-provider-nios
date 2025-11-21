// Create a DTC Monitor SIP with Basic Fields
resource "nios_dtc_monitor_sip" "dtc_monitor_sip_basic_fields" {
  name = "example-monitor-sip"
}

// Create a DTC Monitor SIP with Additional Fields
resource "nios_dtc_monitor_sip" "dtc_monitor_http_additional_fields" {
  name        = "example-monitor-sip-all-fields"
  ciphers     = "DHE-RSA-AES256-SHA"
  client_cert = "dtc:certificate/"
  comment     = "This is a comment"
  extattrs = {
    Site = "location-1"
  }
  interval      = 5
  port          = 80
  timeout       = 30
  request       = "GET /api/health HTTP/1.1\nHost: example.com\nUser-Agent: NIOS-Monitor"
  result        = "CODE_IS"
  result_code   = 400
  retry_down    = 2
  retry_up      = 5
  validate_cert = true
  transport     = "TCP"
}