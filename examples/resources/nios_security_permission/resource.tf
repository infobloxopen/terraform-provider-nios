// Create a View (Required as parent)
resource "nios_dns_view" "example_view" {
  name = "parent_view_for_permission"
}

// Create a Permission with Basic Fields
resource "nios_security_permission" "example_permission_basic_fields" {
  group         = "cloud-api-only"
  permission    = "WRITE"
  resource_type = "HOST"
  object        = nios_dns_view.example_view.ref
}

// Addition of more permissions to same object
resource "nios_security_permission" "example_permission_additional_fields" {
  group         = "cloud-api-only"
  permission    = "READ"
  resource_type = "ZONE"
  object        = nios_dns_view.example_view.ref
}
