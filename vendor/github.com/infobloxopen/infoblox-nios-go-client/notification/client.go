/*
Infoblox NOTIFICATION API

OpenAPI specification for Infoblox NIOS WAPI NOTIFICATION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"github.com/infobloxopen/infoblox-nios-go-client/internal"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
)

const serviceBasePath = "/wapi/v2.13.6"

// APIClient manages communication with the Infoblox NOTIFICATION API 2.13.6
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	NotificationRestEndpointAPI NotificationRestEndpointAPI
	NotificationRestTemplateAPI NotificationRestTemplateAPI
	NotificationRuleAPI NotificationRuleAPI
}

// NewAPIClient creates a new API client.
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithNIOSHostUrl(string) sets the URL for NIOS Portal.
// - WithNIOSUsername(string) sets the Username for the NIOS Portal.
// - WithNIOSPassword(string) sets the Password for the NIOS Portal.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultExtAttrs(map[string]struct{ Value string }) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(serviceBasePath, cfg)

	// API Services
	c.NotificationRestEndpointAPI = (*NotificationRestEndpointAPIService)(&c.Common)
	c.NotificationRestTemplateAPI = (*NotificationRestTemplateAPIService)(&c.Common)
	c.NotificationRuleAPI = (*NotificationRuleAPIService)(&c.Common)

	return c
}