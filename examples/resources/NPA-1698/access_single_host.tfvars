# Test Scenario: Single Host /32
# Named ACL with single host CIDR

name = "acl_single_host"
access_list = [
  {
    struct     = "addressac"
    address    = "10.0.0.100/32"
    permission = "ALLOW"
  }
]
