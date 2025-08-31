// Create Distribution Schedule with basic fields
resource "nios_grid_distributionschedule" "basic_schedule" {
  active = true

  # Note: Replace this with a future UNIX timestamp (must be later than current time)
  start_time = 1756800000
}

// Create Distribution Schedule with additional fields
resource "nios_grid_distributionschedule" "schedule_with_upgrade_groups" {
  active = true

  # Note: Replace this with a future UNIX timestamp (must be later than current time)
  start_time = 1756800000

  upgrade_groups = [
    {
      name = "Default"
      # Note: Replace with a UNIX timestamp later than start_time
      distribution_time = 1756803600
    }
  ]
}
