// Create a Parental Control Subscriber Record with Basic Fields
resource "nios_parentalcontrol_subscriberrecord" "subscriber_record_with_basic_fields" {
  ip_addr       = "10.36.0.154"
  ipsd          = "N/A"
  localid       = "N/A"
  prefix        = 32
  site          = "site1"
  subscriber_id = "IMSI=12345"
}

// Create a Parental Control Subscriber Record with Additional Fields
resource "nios_parentalcontrol_subscriberrecord" "subscriber_record_with_additional_fields" {
  ip_addr       = "10.36.0.155"
  ipsd          = "N/A"
  localid       = "N/A"
  prefix        = 32
  site          = "site1"
  subscriber_id = "IMSI=12345"

  // Additional fields
  accounting_session_id = "Acct-Session-Id=9999732d-34590346"
  alt_ip_addr           = "2123:345:287::6727:22"

  ans0 = "User-Name=JOHN"
  ans1 = "User-Name=JOHN1"
  ans2 = "User-Name=JOHN2"
  ans3 = "User-Name=JOHN3"
  ans4 = "User-Name=JOHN4"
  #   black_list = "[http://a1.com|http://a1.com|smart-link],[a7.com|http://a7.com]"
  bwflag                  = true
  dynamic_category_policy = false
  flags                   = "SB"
  nas_contextual          = "NAS-PORT=1813"
  //op_code = ""
  parental_control_policy  = "00000000000000000000000000020040"
  proxy_all                = true
  subscriber_secure_policy = "FF"
  unknown_category_policy  = false
  #   white_list = "[https://example.com]"
  wpc_category_policy = 1
}


terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.1.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.190"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}