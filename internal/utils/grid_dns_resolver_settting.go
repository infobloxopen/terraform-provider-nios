package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

// ConfigureGridDNSResolver configures grid-level DNS resolver settings
// This function:
// 1. Gets the grid reference using the Grid client
// 2. Updates the grid with DNS resolver settings using the Grid client
// 3. Restarts the grid services using the WAPI function call
func ConfigureGridDNSResolver(ctx context.Context, gridClient *grid.APIClient, dnsResolverSetting *grid.MemberDnsResolverSetting) error {

	gridRef, err := getGridReferenceUsingClient(ctx, gridClient)
	if err != nil {
		return fmt.Errorf("error getting grid reference: %w", err)
	}
	err = updateGridDNSResolverUsingClient(ctx, gridClient, gridRef, dnsResolverSetting)
	if err != nil {
		return fmt.Errorf("error updating grid DNS resolver: %w", err)
	}

	// Extract configuration from grid client for restart services call
	baseURL := gridClient.Cfg.NIOSHostURL
	username := gridClient.Cfg.NIOSUsername
	password := gridClient.Cfg.NIOSPassword

	// Restart grid services
	tflog.Debug(ctx, "Restarting grid services")
	err = restartGridServices(ctx, baseURL, username, password, gridRef)
	if err != nil {
		return fmt.Errorf("error restarting grid services: %w", err)
	}
	tflog.Info(ctx, "Successfully initiated grid services restart")

	tflog.Info(ctx, "Waiting 60 seconds for grid services to stabilize")
	time.Sleep(60 * time.Second)

	return nil
}

// getGridReferenceUsingClient retrieves the grid reference using the Grid client
func getGridReferenceUsingClient(ctx context.Context, gridClient *grid.APIClient) (string, error) {
	apiRes, _, err := gridClient.GridAPI.List(ctx).ReturnAsObject(1).Execute()
	if err != nil {
		return "", fmt.Errorf("error listing grid objects: %w", err)
	}

	res := apiRes.ListGridResponseObject.GetResult()
	if len(res) > 0 {
		if res[0].Ref == nil {
			return "", fmt.Errorf("grid reference is nil in response")
		}
		return *res[0].Ref, nil
	}
	return "", fmt.Errorf("no grid reference found in response")
}

// updateGridDNSResolverUsingClient updates the grid with DNS resolver settings using the Grid client
func updateGridDNSResolverUsingClient(ctx context.Context, gridClient *grid.APIClient, gridRef string, dnsResolverSetting *grid.MemberDnsResolverSetting) error {
	gridDnsResolverSetting := grid.GridDnsResolverSetting{
		Resolvers:     dnsResolverSetting.Resolvers,
		SearchDomains: dnsResolverSetting.SearchDomains,
	}

	updateReq := grid.NewGrid()
	updateReq.DnsResolverSetting = &gridDnsResolverSetting

	tflog.Debug(ctx, fmt.Sprintf("Updating grid %s with DNS resolver settings", gridRef))

	_, _, err := gridClient.GridAPI.Update(ctx, ExtractResourceRef(gridRef)).Grid(*updateReq).Execute()
	if err != nil {
		return fmt.Errorf("error updating grid with DNS resolver settings: %w", err)
	}

	return nil
}

// restartGridServices restarts the grid services using the WAPI function call
func restartGridServices(ctx context.Context, baseURL, username, password, gridRef string) error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	restartURL := fmt.Sprintf("%s/wapi/v2.14/%s?_function=restartservices", baseURL, gridRef)

	req, err := http.NewRequestWithContext(ctx, "POST", restartURL, nil)
	if err != nil {
		return fmt.Errorf("error creating restart services request: %w", err)
	}

	req.SetBasicAuth(username, password)

	tflog.Debug(ctx, fmt.Sprintf("Making restart services request to: %s", restartURL))
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making restart services request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("restart services request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
