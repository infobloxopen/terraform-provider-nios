// Retrieve a specific DTC Pool by filters 
data "nios_dtc_pool" "get_record_using_filters" {
  filters = {
    name = "dtc_pool"
  }
}

// Retrieve specific DTC Pool using Extensible Attributes
data "nios_dtc_pool" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Pools
data "nios_dtc_pool" "get_all_pools" {}
