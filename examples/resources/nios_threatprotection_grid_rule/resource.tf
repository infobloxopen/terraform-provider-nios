// Manage a Threat Protection Grid Rule with Basic Fields
resource "nios_threatprotection_grid_rule" "threatprotection_grid_rule_basic" {
  template = ""
}

// Manage a Threat Protection Grid Rule with Additional Fields
resource "nios_threatprotection_grid_rule" "threatprotection_grid_rule_with_additional_fields" {
  template = ""

  // Additional Fields
  comment = "Threat Protection Grid Rule created by Terraform"
  config = {
    action       = "ALERT"
    log_severity = "CRITICAL"
    params = [
      {
        name  = "param1"
        value = "value1"
      }
    ]
  }
  disabled = false
}
