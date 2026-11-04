action "nios_restart_services" "dns" {
  config {
    services       = ["DNS"]
    restart_option = "FORCE_RESTART"
    members        = ["infoblox.member2"]
    user_name      = "admin"
    mode           = "SEQUENTIAL"
  }
}

# Restart the DNS Service once the Grid Primary is added to Zone Auth
resource "nios_dns_zone_auth" "create_zone5" {
  // Basic Fields
  fqdn = "example_restart_services.com"
  view = "default"

  grid_primary = [
    {
      name = "infoblox.member2",
    }
  ]

  comment = "Zone Auth Restart Services Example"
  extattrs = {
    Site = "location-1"
  }

  lifecycle {
    action_trigger {
      events  = [after_create, after_update]
      actions = [action.nios_restart_services.dns]
    }
  }
}
