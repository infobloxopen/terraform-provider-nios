terraform {
    required_providers {
        nios = {
            source  = "infoblox-cto/nios"
            version = "1.0.0"
        }
    }
}

provider "nios" {
    nios_auth="admin:infoblox"
    nios_host_url="https://10.197.81.146"
    }
