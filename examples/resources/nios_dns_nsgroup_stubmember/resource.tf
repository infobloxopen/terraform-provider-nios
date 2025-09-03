terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.81.186"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

// Create a NS Group Stubmember with Basic Fields
resource "nios_dns_nsgroup_stubmember" "nsgroup_stubmember_with_basic_fields"{
    name = "stubmember1"
    stub_members = [
        {
            name = "member.com"
        }
    ]
}