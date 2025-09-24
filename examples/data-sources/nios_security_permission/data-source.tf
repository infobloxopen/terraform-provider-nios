// Retrieve a specific Permission by filters
data "nios_security_permission" "get_permission_using_filters" {
  filters = {
    group         = "cloud-api-only"
    permission    = "WRITE"
    resource_type = "HOST"
  }
}

// Retrieve all Permissions
data "nios_security_permission" "get_all_permissions" {}
