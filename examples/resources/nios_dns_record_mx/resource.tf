// Create MX Record with Basic Fields
resource "nios_dns_record_mx" "record1" {
  name           = "mx_record.example.com"
  mail_exchanger = "mail.example.com"
  preference     = 10
  view           = "default"
}

// Create MX Record with additional fields
resource "nios_dns_record_mx" "record2" {
  name           = "example.com"
  mail_exchanger = "mail1.example.com"
  preference     = 20
  view           = "default"
  use_ttl        = true
  ttl            = 3600
  comment        = "Example MX Record"
  extattrs = {
    Site = "location-1"
  }
}
