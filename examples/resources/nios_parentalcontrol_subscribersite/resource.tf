// Create a Parental Control Subscriber Site with Basic Fields
resource "nios_parentalcontrol_subscribersite" "subscriber_site_with_basic_fields" {
  name = "example_subscriber_site-1"
}

// Create a Parental Control Subscriber Site with Additional Fields
resource "nios_parentalcontrol_subscribersite" "subscriber_site_with_additional_fields" {
  name = "example_subscriber_site_2"

  // Additional Fields
  abss = [
    {
      blocking_policy = "policy1",
      ip_address      = "12.13.12.12"
    }
  ]
  block_size                   = 24
  blocking_ipv4_vip1           = "12.13.11.11"
  blocking_ipv4_vip2           = "14.12.12.13"
  blocking_ipv6_vip1           = "2001:db8::1"
  blocking_ipv6_vip2           = "2001:db8::2"
  comment                      = "Example Subscriber Site"
  dca_sub_bw_list              = true
  dca_sub_query_count          = true
  enable_global_allow_list_rpz = false
  enable_rpz_filtering_bypass  = false
  first_port                   = 10312
  global_allow_list_rpz        = 34
  maximum_subscribers          = 100000
  msps = [
    {
      ip_address : "12.12.14.1"
    }
  ]
  nas_port = 56
  nas_gateways = [
    {
      "ip_address" = "12.1.1.1",
      "name"       = "nas_gateway_1",
      "send_ack" : false,
      "shared_secret" : "secret123"
    }
  ]
  members = [
    {
      "name" = "infoblox.localdomain"
    }
  ]
  api_members = [
    {
      "name" = "infoblox.localdomain"
    }
  ]
  proxy_rpz_passthru = false
  spms = [
    {
      ip_address : "12.13.14.1"
    }
  ]
  stop_anycast               = true
  strict_nat                 = false
  subscriber_collection_type = "API"
  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
