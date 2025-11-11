
// Create a Shared MX Record with Basic Fields
resource "nios_dns_sharedrecord_mx" "sharedrecord_mx_basic_fields" {
    mail_exchanger = "mail.example.com"
    name = "sharedmx_record"
    preference = 10
    shared_record_group = "example-sharedrecordgroup"
}

// Create a Shared MX Record with Additional Fields
resource "nios_dns_sharedrecord_mx" "sharedrecord_mx_additional_fields" {
    mail_exchanger = "mail.example.com"
    name = "sharedmx_record_additional_fields"
    preference = 20
    shared_record_group = "example-sharedrecordgroup"
    comment = "Example Shared MX Record"
    disable = true
    extattrs = {
        Site = "location-1"
    }
    use_ttl = true
    ttl = 7200
}
