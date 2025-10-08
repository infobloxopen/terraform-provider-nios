// Create an Admin Group with Basic Fields
resource "nios_security_admin_group" "admin_group_with_basic_fields" {
  name = "example_admin_group"
}

// Create an Admin Group with Additional Fields with user_access as IP address
resource "nios_security_admin_group" "admin_group_with_additional_fields" {
  name          = "example_admin_group2"
  access_method = ["GUI", "API", "CLI", "TAXII"]
  admin_set_commands = {
    set_bfd = true
    set_bfg = false
  }
  admin_show_commands = {
    show_bfd = true
    show_bfg = false
  }
  admin_toplevel_commands = {
    iostat  = true
    netstat = true
  }
  cloud_set_commands = {
    set_cloud_services_portal = true
  }
  database_set_commands = {
    set_db_rollover = true
    set_db_snapshot = false
  }
  dhcp_set_commands = {
    set_overload_bootp = true
  }
  dns_show_commands = {
    show_dns       = true
    show_dns_accel = false
  }
  dns_set_commands = {
    set_dns       = true
    set_dns_accel = false
  }
  dns_toplevel_commands = {
    ddns_delete = true
    ddns_add    = false
  }
  docker_set_commands = {
    set_docker_bridge = true
  }
  docker_show_commands = {
    show_docker_bridge = true
  }
  email_addresses = ["abd@info.com", "xyz@example.com"]
  grid_set_commands = {
    set_dscp       = true
    set_membership = false
  }
  grid_show_commands = {
    show_token = true
  }
  inactivity_lockout_setting = {
    inactive_days                        = 20
    reactivate_via_remote_console_enable = true
    reminder_days                        = 10
  }
  use_account_inactivity_lockout_enable = true
  password_setting = {
    expire_days   = 90
    expire_enable = true
    reminder_days = 15
  }
  use_password_setting = true
  roles                = ["DHCP Admin", "DNS Admin", "Grid Admin"]
  saml_setting = {
    auto_create_user          = true
    persist_auto_created_user = false
  }
  superuser = true
  user_access = [
    {
      address    = "12.12.1.1"
      permission = "ALLOW"
    }
  ]
}

// Create Named Access Control Lists (ACLs) ( required as parent )
resource "nios_acl_namedacl" "namedacl_with_basic_fields" {
  name = "example-named-acl"
  access_list = [
    {
      struct     = "addressac"
      address    = "10.0.0.1"
      permission = "ALLOW"
    },
    {
      struct     = "addressac"
      address    = "10.0.0.2"
      permission = "ALLOW"
    }
  ]
}

// Create an Admin Group with Additional Fields with user_access as Named ACL
resource "nios_security_admin_group" "admin_group_with_additional_fields2" {
  name    = "example_admin_group3"
  comment = "Example Admin Group with additional fields"
  disable = true
  extattrs = {
    Site = "location-1"
  }
  cloud_show_commands = {
    show_cloud_services_portal = true
  }
  database_show_commands = {
    show_backup                   = true
    show_database_transfer_status = false
  }
  dhcp_show_commands = {
    show_log_txn_id               = false
    show_dhcpd_recv_sock_buf_size = true
  }
  licensing_set_commands = {
    set_license                 = true
    set_reporting_reset_license = false
  }
  licensing_show_commands = {
    show_license                = true
    show_license_pool_container = false
  }
  lockout_setting = {
    failed_lockout_duration                         = 30
    never_unlock_user                               = false
    sequential_attempts                             = 3
    enable_sequential_failed_login_attempts_lockout = true
  }
  use_lockout_setting = true
  machine_control_toplevel_commands = {
    reboot   = true
    restart  = false
    reset    = true
    shutdown = false
  }
  networking_set_commands = {
    set_connection_limit = true
    set_network          = false
  }
  networking_show_commands = {
    show_connection_limit = true
    show_connections      = false
  }
  security_set_commands = {
    set_adp     = true
    set_cc_mode = false
  }
  security_show_commands = {
    show_cc_mode                 = true
    show_certificate_auth_admins = false
  }
  trouble_shooting_toplevel_commands = {
    console = true
    dig     = false
    ping    = true
  }
  disable_concurrent_login     = false
  use_disable_concurrent_login = true
  user_access = [
    {
      ref = nios_acl_namedacl.namedacl_with_basic_fields.ref
    }
  ]
  depends_on = [nios_acl_namedacl.namedacl_with_basic_fields]
}
