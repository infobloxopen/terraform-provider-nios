// Retrieve a specific DTC LBDN record by filters
data "nios_dtc_lbdn" "get_record_with_filter" {
  filters = {
    name = "example_lbdn"
  }
}

// Retrieve specific DTC LBDN records using Extensible Attributes
data "nios_dtc_lbdn" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "Spain"
  }
}

// Retrieve all DTC LBDN records
data "nios_dtc_lbdn" "get_all_records" {}
