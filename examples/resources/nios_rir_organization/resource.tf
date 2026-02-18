// Manage an RIR Organization with Basic Fields
resource "nios_rir_organization" "rir_organization_basic" {
  id           = "ORG-CR07-IB"
  maintainer   = "infoblox"
  name         = "example_rir_organization"
  password     = "example-pass"
  sender_email = "support@infoblox.com"
  extattrs = {
    "RIPE Admin Contact" : "ib-contact",
    "RIPE Country" : "United Kingdom (GB)",
    "RIPE Technical Contact" : "EG123-IB",
    "RIPE Email" : "support@infoblox.com",
  }
}

// Manage an RIR Organization with Additional Fields
resource "nios_rir_organization" "rir_organization_with_additional_fields" {
  id           = "ORG-BB99-IB"
  maintainer   = "nios"
  name         = "example_rir_organization_additional"
  password     = "examplePass123"
  sender_email = "support@infoblox.com"
  extattrs = {
    "RIPE Admin Contact" : "ib-contact",
    "RIPE Country" : "United Kingdom (GB)",
    "RIPE Technical Contact" : "EG567-IB",
    "RIPE Email" : "support@infoblox.com",

    // Additional Extensible Attributes
    "RIPE Remarks" : "Example RIR Organization"
    "RIPE Organization Type" : "IANA"
    "RIPE Notify" : "support@infoblox.com"
  }

  // Additional Fields
  rir = "RIPE"
}
