terraform {
    required_providers {
        nios = {
            source  = "infoblox-cto/nios"
            version = "1.0.0"
        }
    }
}

provider "nios" {
    nios_auth="username:password"
    nios_host_url="<Nios Host URL>"
}