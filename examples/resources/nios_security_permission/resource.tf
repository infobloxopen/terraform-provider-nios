//Create a View (Required as parent )
resource "nios_dns_view" "example_view" {
  name = "View_for_permission"
}

// Create a Permission with Basic Fields
resource "nios_security_permission" "example_permission_basic_fields" {
  group         = "cloud-api-only"
  permission    = "WRITE"
  resource_type = "HOST"
  object        = nios_dns_view.example_view.ref
}
