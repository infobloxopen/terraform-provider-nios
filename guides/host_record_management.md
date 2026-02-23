# Host Record Management

The `nios_ip_allocation` resource allocates a new IP address from an existing NIOS network and manages the corresponding DNS-related settings. It creates a Host Record in NIOS with either an IPv4 address, an IPv6 address, or both. The IP can be allocated statically (by specifying the address) or dynamically (as the next available address from a network). Once allocated, the address is marked as **used** in NIOS.

The `nios_ip_association` resource manages DHCP-related properties of the Host Record created by `nios_ip_allocation`. It attaches instance network identifiers (MAC for IPv4, DUID for IPv6) and DHCP configuration so the allocated IP can be associated with a VM or instance.

Host Record serves as the backend for the following operations:

- Allocation and deallocation of an IP address from a network (`nios_ip_allocation`)
- Association and disassociation of an IP address with a VM (`nios_ip_association`)

---

## Notes

- `nios_ip_allocation` **owns the Host Record lifecycle**. Create the allocation first - it creates the Host Record and DNS entries.
- `nios_ip_association` is **read & update only** and must reference an existing allocation via `ref`. Do **not** create the association before the allocation exists.
- Do **not** destroy the association resource directly - destroy the allocation and the provider will remove the associated Host Record and DHCP settings.
- Each allocation supports **at most one IPv4 and one IPv6** address. You may have one of each (dual-stack), but multiple addresses of the same family are not supported.
- Use `nios_ip_association` to attach instance identifiers (MAC / DUID) and DHCP flags only; allocation handles DNS and address allocation.

### Recommended workflow
1. Create `nios_ip_allocation` (allocates IP + manages DNS).
2. Create `nios_ip_association` referencing allocation (`ref = nios_ip_allocation.<name>.ref`) to attach MAC/DUID and DHCP settings.

### Destroy order
- To remove a Host Record and its DHCP association:
  - Destroy `nios_ip_allocation` (this will remove the Host Record and the association).
  - Do NOT destroy the `nios_ip_association` first.

---

## Examples

For additional Terraform configurations, see the [Terraform resource examples](../examples/resources/nios_dns_record_host/).

> **Note:** It’s recommended to reference the allocation (`ref = nios_ip_allocation.<name>.ref`) in the association to create an implicit dependency. This ensures Terraform builds the correct dependency graph and executes the allocation before the association, avoiding race conditions in parallel execution.

### Static IPv4 Host Record with MAC Address

```terraform
# Allocate a static IPv4 address with DNS configuration
resource "nios_ip_allocation" "allocation_static" {
  name              = "host1.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.110"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Associate MAC address without enabling DHCP
resource "nios_ip_association" "association_static" {
  ref                = nios_ip_allocation.allocation_static.ref
  mac                = "12:00:43:fe:9a:8c"
  configure_for_dhcp = false
}
```

### Dual-Stack Host Record with DHCP Enabled

```hcl
# Allocate both IPv4 and IPv6 addresses
resource "nios_ip_allocation" "allocation_dual_stack" {
  name              = "host2.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.112"
    }
  ]
  ipv6addrs = [
    {
      ipv6addr = "2002:1f93::12:2"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Associate MAC and DUID with DHCP enabled
resource "nios_ip_association" "association2_dual_stack" {
  ref                = nios_ip_allocation.allocation_dual_stack.ref
  mac                = "12:43:fd:ba:9c:c9"
  duid               = "00:01:5f:3a:1b:2c:12:34:56:78:9a:bc"
  match_client       = "DUID"
  configure_for_dhcp = true
}
```

### Dynamic IP Allocation Using Next Available Address

```hcl
# Dynamically allocate next available IP from network
resource "nios_ip_allocation" "allocation_dynamic" {
  name              = "host3.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      func_call = {
        attribute_name  = "ipv4addr"
        object_function = "next_available_ip"
        result_field    = "ips"
        object          = "network"
        object_parameters = {
          network      = "10.10.0.0/16"
          network_view = "default"
        }
        parameters = {
          exclude = jsonencode(["10.10.0.1", "10.10.0.2"])
        }
      }
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Associate MAC address without enabling DHCP
resource "nios_ip_association" "association_dynamic" {
  ref                = nios_ip_allocation.allocation_dynamic.ref
  mac                = "12:00:43:fe:9a:8d"
  configure_for_dhcp = false
}
```