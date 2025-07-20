// Create Record SRV with Basic Fields
resource "nios_dns_record_srv" "create_record" {
    name = "example-srv-record.example.com"
    target = "example.target.com"
    port = 80
    priority = 4
    weight = 50
    
    // Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}

// Create Record SRV with additional fields
resource "nios_dns_record_srv" "create_with_additional_config" {
    name = "example-srv-record-with-config.example.com"
    target = "example_updated.target.com"
    port = 8080
    priority = 2
    weight = 100
    
    // Additional Fields
    view     = "default"
    use_ttl  = true
    ttl      = 10
    creator = "DYNAMIC"
    comment  = "Example SRV record"
    
    // Extensible Attributes
    extattrs = {
        Site = "location-2"
    }
}
