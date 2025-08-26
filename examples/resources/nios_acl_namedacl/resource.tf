// Create Named Access Control Lists (ACLs) with basic fields
resource "nios_acl_namedacl" "base_namedacl" {
  name    = "base-acl-template"
  comment = "Base ACL structure created for future assignment of access control entries"

  extattrs = {
    Site = "location-1"
  }
}

// Create Named Access Control Lists (ACLs) with additional fields
resource "nios_acl_namedacl" "namedacl_with_additional_config" {
  name    = "dev-network-acl"
  comment = "ACL to allow/deny access to specific dev network resources"

  // ACL Entries
  acl = [
    {
      struct     = "addressac"
      address    = "10.0.0.1"
      permission = "ALLOW"
    },
    {
      struct     = "addressac"
      address    = "10.0.0.2"
      permission = "DENY"
    }
  ]

  extattrs = {
    Site = "location-2"
  }
}
