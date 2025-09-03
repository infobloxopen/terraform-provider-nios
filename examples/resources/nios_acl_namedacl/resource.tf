// Create Named Access Control Lists (ACLs) with Basic Fields
resource "nios_acl_namedacl" "namedacl_with_basic_fields" {
  name    = "base-acl-template"
  comment = "Base ACL structure created for future assignment of access control entries"

  extattrs = {
    Site = "location-1"
  }
}

// Create Named Access Control Lists (ACLs) with Additional Fields
resource "nios_acl_namedacl" "namedacl_with_additional_config" {
  name    = "example-network-acl"
  comment = "ACL to allow/deny access to specific dev network resources"

  // ACL Entries
  access_list = [
    {
      struct     = "addressac"
      address    = "10.0.0.1"
      permission = "ALLOW"
    },
    {
      struct     = "addressac"
      address    = "10.0.0.2"
      permission = "DENY"
    },
    {
      struct        = "tsigac"
      tsig_key      = "X4oRe92t54I+T98NdQpV2w=="
      tsig_key_name = "example-tsig-key"
      tsig_key_alg  = "HMAC-SHA256"
    }
  ]

  extattrs = {
    Site = "location-2"
  }
}
