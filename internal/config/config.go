package config

var proxySearch string

// SetProxySearch sets the ProxySearch value
func SetProxySearch(ps string) {
	proxySearch = ps
	if proxySearch == "" {
		proxySearch = "LOCAL"
	}
}

// GetProxySearch gets the ProxySearch value
func GetProxySearch() string {
	return proxySearch
}
