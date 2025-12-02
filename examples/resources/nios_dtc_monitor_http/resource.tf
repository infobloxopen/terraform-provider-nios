// Create a DTC Monitor HTTP with Basic Fields
resource "nios_dtc_monitor_http" "dtc_monitor_http_basic_fields" {
  name = "example-monitor-http"
}

// Create a DTC Monitor HTTP with Additional Fields
resource "nios_dtc_monitor_http" "dtc_monitor_http_additional_fields" {
  name                  = "example-monitor-http-all-fields"
  ciphers               = "DHE-RSA-AES256-SHA"
  client_cert           = "dtc:certificate/example"
  comment               = "Example DTC Monitor HTTP"
  content_check         = "EXTRACT"
  content_check_input   = "BODY"
  content_check_op      = "EQ"
  content_check_regex   = "Load: ([0-9]+)"
  content_extract_group = 1
  content_extract_type  = "STRING"
  content_extract_value = "SUCCESS"
  enable_sni            = false
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
  secure        = false
  validate_cert = false
}
