// Create an Attribute Value Pair with Basic Fields
resource "nios_parentalcontrol_avp" "attribute_value_pair_basic" {
  name       = "avp_basic"
  type       = 42
  value_type = "BYTE"
}

// Create an Attribute Value Pair with Additional Fields
resource "nios_parentalcontrol_avp" "attribute_value_pair_additional_fields" {
  name          = "avp_addional_fields"
  type          = 37
  value_type    = "INTEGER64"
  comment       = "Example AVP"
  domain_types  = ["SUBS_ID", "IP_SPACE_DIS"]
  is_restricted = true
  vendor_id     = 1234
  vendor_type   = 230
}

terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "0.0.1"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.83.204"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}