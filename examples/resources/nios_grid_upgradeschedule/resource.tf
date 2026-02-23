// Update Upgrade Schedule with Basic Fields
resource "nios_grid_upgradeschedule" "basic_schedule" {
  active     = true
  start_time = "2026-10-09T20:30:00"
}

// Update Upgrade Schedule with Additional Fields
resource "nios_grid_upgradeschedule" "schedule_with_upgrade_groups" {
  active     = true
  start_time = "2026-10-09T20:00:00"
  upgrade_groups = [
    {
      name         = "Default"
      upgrade_time = "2026-10-09T22:30:00"
    }
  ]
}
