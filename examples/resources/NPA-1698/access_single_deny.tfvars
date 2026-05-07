# Test Scenario: Single Address Deny
# Named ACL with single address deny entry

name = "acl_single_deny"
access_list = [
  {
    struct     = "addressac"
    address    = "192.168.1.100"
    permission = "DENY"
  }
]
