//Create a DTC Server with basic fields
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

//Create a DTC Server with additional fields. 
resource "nios_dtc_server" "create_with_additional_fields" {
  name                    = "example-server"
  host                    = "2.3.3.4"
  auto_create_host_record = true
  comment                 = "create server"
  disable                 = false
  extattrs = {
    Site = "location-1"
  }
  monitors = [
    {
      monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"
      host    = "3.23.23.3"
    },
    {
      monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      host    = "3.3.3.2"
    }
  ]
  sni_hostname     = "server-sni"
  use_sni_hostname = true
}
