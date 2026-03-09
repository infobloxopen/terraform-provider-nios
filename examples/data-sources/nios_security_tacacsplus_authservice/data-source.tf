// Retrieve a specific Tacacsplus Authservice by filters
data "nios_security_tacacsplus_authservice" "get_tacacsplus_authservice_using_filters" {
  filters = {
    name = "tacacsplus_authservice1"
  }
}

// Retrieve all Tacacsplus Authservices
data "nios_security_tacacsplus_authservice" "get_all_tacacsplus_authservices" {}
