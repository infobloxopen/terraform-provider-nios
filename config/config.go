package config

var suppressComputedPlan bool

// SetSuppressComputedPlan sets the suppressComputedPlan flag.
// If not provided, defaults to false.
func SetSuppressComputedPlan(enabled bool) {
	suppressComputedPlan = enabled
}

// GetSuppressComputedPlan returns the suppressComputedPlan flag value.
func GetSuppressComputedPlan() bool {
	return suppressComputedPlan
}
