// Retrieve a specific Permission by filters
data "nios_security_permission" "example_permission" {
  filters = {
    group         = "cloud-api-only"
    permission    = "WRITE"
    resource_type = "HOST"
  }
}
