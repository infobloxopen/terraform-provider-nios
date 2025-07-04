//create a DTC Server with minimal parameters 
resource "nios_dtc_server" "dtc_server" {
  name = "test-server"
  host = "2.3.3.4"
}

//create a DTC Server with maximal parameters 
resource "nios_dtc_server" "dtc_server_maximal_parameters" {
  name                    = "test-server"
  host                    = "2.3.3.4"
  auto_create_host_record = true
  comment                 = "create server"
  disable                 = false
  extattrs = {
    Site = "Siteblr"
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
