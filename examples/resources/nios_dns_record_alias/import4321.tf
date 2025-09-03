# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "record:alias/ZG5zLmFsaWFzX3JlY29yZCQuX2RlZmF1bHQuY29tLmluZm8uaW1wb3J0LWFsaWFzLk5BUFRS:import-alias.info.com/default"
resource "nios_dns_record_alias" "import4321" {
  comment = null
  creator = "STATIC"
  disable = false
  extattrs = {
    "IB Discovery Owned" = "aaddd"
    Site                 = "Blr"
  }
  name        = "import-alias.info.com"
  target_name = "aa.bb.com"
  target_type = "NAPTR"
  ttl         = 100
  use_ttl     = false
  view        = "default"
}
