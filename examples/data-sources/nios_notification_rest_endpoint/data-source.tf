// Retrieve a specific Notification REST Endpoint by filters
data "nios_notification_rest_endpoint" "get_notification_rest_endpoint_using_filter" {
  filters = {
    name = "example-notification-rest-endpoint-1"
  }
}

// Retrieve specific Notification REST Endpoints using Extensible Attributes
data "nios_notification_rest_endpoint" "get_notification_rest_endpoint_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Notification REST Endpoints
data "nios_notification_rest_endpoint" "get_all_notification_rest_endpoints" {}
