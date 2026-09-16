variable "dhcp_ranges" {
  description = "DHCP ranges that should be applied before restarting the DHCP service."
  type = map(object({
    start_addr = string
    end_addr   = string
  }))
  default = {
    primary = {
      start_addr = "10.0.0.170"
      end_addr   = "10.0.0.180"
    }
    secondary = {
      start_addr = "10.0.0.190"
      end_addr   = "10.0.0.200"
    }
  }
}

resource "nios_ipam_network" "example" {
  network      = "10.0.0.0/24"
  network_view = "default"
}

resource "nios_dhcp_range" "example" {
  for_each = var.dhcp_ranges

  start_addr = each.value.start_addr
  end_addr   = each.value.end_addr

  depends_on = [nios_ipam_network.example]
}

action "nios_restart_services" "dhcp" {
  config {
    services       = ["DHCP"]
    restart_option = "RESTART_IF_NEEDED"
  }
}

# Update once after the complete set of DHCP ranges has changed. The dependency
# ensures that Terraform invokes the action only after the ranges are ready.
resource "terraform_data" "dhcp_restart" {
  input = sha256(jsonencode(var.dhcp_ranges))

  depends_on = [nios_dhcp_range.example]

  lifecycle {
    action_trigger {
      events  = [after_create, after_update]
      actions = [action.nios_restart_services.dhcp]
    }
  }
}
