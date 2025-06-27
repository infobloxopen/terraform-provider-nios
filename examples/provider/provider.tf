terraform {
    required_providers {
        nios = {
            source  = "infoblox-cto/nios"
            version = "1.0.0"
        }
    }
}

provider "nios" {
    nios_host_url="<NIOS_HOST_URL>"
    nios_username="<NIOS_USERNAME>"
    nios_password="<NIOS_PASSWORD>"
}
