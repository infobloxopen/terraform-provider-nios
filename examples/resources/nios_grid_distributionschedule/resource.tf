// Update Distribution Schedule with basic fields
resource "nios_grid_distributionschedule" "basic_schedule" {
  active     = true
  start_time = "2026-09-09T20:30:00"
}

// Update Distribution Schedule with additional fields
resource "nios_grid_distributionschedule" "schedule_with_upgrade_groups" {
  active     = true
  start_time = "2026-09-09T20:00:00"
  upgrade_groups = [
    {
      name              = "Default"
      distribution_time = "2026-09-09T22:30:00"
    }
  ]
}
