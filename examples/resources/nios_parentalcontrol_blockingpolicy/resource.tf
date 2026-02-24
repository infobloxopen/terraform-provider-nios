// Create a Parental Control Blocking Policy
resource "nios_parentalcontrol_blockingpolicy" "parentalcontrol_blockingpolicy" {
  name  = "example-blockingpolicy-2"
  value = "00000000000000000000000000009002"
}
