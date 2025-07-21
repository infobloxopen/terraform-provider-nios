terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.248"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

resource "nios_dhcp_range" "test"{
    start_addr = "10.0.0.170"
    end_addr   = "10.0.0.180"
    # exclude = [
    #     {
    #         start_address = "10.0.0.42"
    #         end_address   = "10.0.0.47"
    #     }
    # ]
    email_list =["chaithra@infoblox.com","example@infoblox.com"]
    use_email_list = true
}