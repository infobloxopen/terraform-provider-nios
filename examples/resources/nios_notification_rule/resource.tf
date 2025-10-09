// Create a Notification Rest Endpoint with Basic Fields ( Required as parent for Notification Rule )
resource "nios_notification_rest_endpoint" "notification_rest_parent {
  name                 = "example-notification-rest-endpoint-1"
  outbound_member_type = "GM"
  uri                  = "https://example.com"
}


// Create a Notification Rule with Basic Fields (Event Type : DHCP_LEASES)
resource "nios_notification_rule" "notification_rule_with_basic_fields" {
  name                = "example_notification_rule"
  notification_target = nios_notification_rest_endpoint.notification_rest_parent.ref
  event_type          = "DHCP_LEASES"
  notification_action = "RESTAPI_TEMPLATE_INSTANCE"

  expression_list = [
    {
      op       = "AND"
      op1_type = "LIST"
    },
    {
      op       = "EQ"
      op1      = "DHCP_LEASE_STATE"
      op1_type = "FIELD"
      op2      = "DHCP_LEASE_STATE_ACTIVE"
      op2_type = "STRING"
    },
    {
      op = "ENDLIST"
    }
  ]
  template_instance = {
    template = "DHCP_Lease"
  }
}

// Create a Notification Rule with Additional Fields
resource "nios_notification_rule" "notification_rule_with_additional_fields" {
  name                = "example_notification_rule-1"
  notification_target = nios_notification_rest_endpoint.notification_rest_parent.ref
  event_type          = "DHCP_LEASES"
  notification_action = "RESTAPI_TEMPLATE_INSTANCE"

  expression_list = [
    {
      op       = "AND"
      op1_type = "LIST"
    },
    {
      op       = "EQ"
      op1      = "DHCP_LEASE_STATE"
      op1_type = "FIELD"
      op2      = "DHCP_LEASE_STATE_ACTIVE"
      op2_type = "STRING"
    },
    {
      op = "ENDLIST"
    }
  ]
  template_instance = {
    template = "DHCP_Lease"
  }

  // Additional Fields
  comment = "Example Notification Rule"
  disable = true
}

// Create a Notification Rule with Basic Fields (Event Type : DNS_RPZ)
resource "nios_notification_rule" "notification_rule_with_dns_rpz" {
  name                = "example-notification-rule-2"
  notification_target = "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
  event_type          = "DNS_RPZ"
  notification_action = "RESTAPI_TEMPLATE_INSTANCE"

  expression_list = [
    {
      op       = "AND"
      op1_type = "LIST"
    },
    {
      op       = "EQ"
      op1      = "DNS_RPZ_TYPE"
      op1_type = "FIELD"
      op2      = "DNS_RPZ_TYPE_IP"
      op2_type = "STRING"
    },
    {
      op = "ENDLIST"
    }
  ]
  template_instance = {
    template = "syslog_action_template"
  }
}

// Create a Notification Rule with Basic Fields (Event Type : IPAM)
resource "nios_notification_rule" "notification_rule_with_ipam_event_type" {
  name                = "example-notification-rule-3"
  notification_target = "pxgrid:endpoint/b25lLmVuZHBvaW50JDU:cisco"
  event_type          = "IPAM"
  notification_action = "RESTAPI_TEMPLATE_INSTANCE"
  template_instance = {
    template = "IPAM_PxgridEvent"
  }
}
