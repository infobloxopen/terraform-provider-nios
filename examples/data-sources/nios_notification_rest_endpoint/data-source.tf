// Retrieve a specific Notification REST Endpoint by filters
data "nios_notification_rest_endpoint" "get_notification_rest_endpoint_using_filter" {
  filters = {
    name = nios_notification_rest_endpoint.test.name
  }
}

// Retrieve specific Notification REST Endpoints using Extensible Attributes
data "nios_notification_rest_endpoint" "get_notification_rest_endpoint_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Notification REST Endpoints
data "nios_notification_rest_endpoint" "get_all_notification_rest_endpoints" {}
