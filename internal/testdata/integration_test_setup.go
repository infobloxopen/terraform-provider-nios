package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

// UploadInitResponse holds the fields returned by the NIOS uploadinit fileop endpoint.
type UploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// UploadInit calls the NIOS fileop uploadinit endpoint and returns the upload
// token and URL needed for subsequent file-upload operations.
func UploadInit(host, wapiVer, username, password string) (*UploadInitResponse, error) {
	url := fmt.Sprintf("%s/wapi/%s/fileop?_function=uploadinit", host, wapiVer)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString("{}"))
	if err != nil {
		return nil, fmt.Errorf("uploadinit: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	fmt.Print(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("uploadinit: execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("uploadinit: unexpected status %s", resp.Status)
	}

	var result UploadInitResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("uploadinit: decode response: %w", err)
	}

	return &result, nil
}

// UploadCertificateRequest holds the body fields for the uploadcertificate fileop call.
type UploadCertificateRequest struct {
	CertificateUsage string `json:"certificate_usage"`
	Member           string `json:"member"`
	Token            string `json:"token"`
}

type UploadFileRequest struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func UploadFile(host, wapiVer, username, password, url, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("uploadfile: open file %q: %w", filePath, err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	formFile, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("uploadfile: create form file: %w", err)
	}

	if _, err := io.Copy(formFile, file); err != nil {
		return fmt.Errorf("uploadfile: copy file contents: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("uploadfile: close multipart writer: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, &requestBody)
	if err != nil {
		return fmt.Errorf("uploadfile: create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploadfile: execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploadfile: unexpected status %s", resp.Status)
	}

	return nil

}

// UploadCertificate uses the token from UploadInit to upload a certificate to NIOS.
// It mirrors:
//
//	curl -k -u admin:infoblox -X POST -H "Content-Type: application/json" \
//	  https://<host>/wapi/<ver>/fileop?_function=uploadcertificate \
//	  -d '{"certificate_usage":"EAP_CA","member":"infoblox.localdomain","token":"<token>"}'
func UploadCertificate(host, wapiVer, username, password, certificateUsage, member, token string) error {
	endpoint := fmt.Sprintf("%s/wapi/%s/fileop?_function=uploadcertificate", host, wapiVer)

	cleanToken := strings.TrimSpace(token)

	body := UploadCertificateRequest{
		CertificateUsage: certificateUsage,
		Member:           member,
		Token:            cleanToken,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("uploadcertificate: marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("uploadcertificate: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	fmt.Print(req)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploadcertificate: execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploadcertificate: unexpected status %s", resp.Status)
	}

	return nil
}

// CACertificate represents a single CA certificate object from the NIOS WAPI.
type CACertificate struct {
	Ref               string `json:"_ref"`
	DistinguishedName string `json:"distinguished_name"`
	Issuer            string `json:"issuer"`
	Serial            string `json:"serial"`
	ValidNotAfter     int64  `json:"valid_not_after"`
	ValidNotBefore    int64  `json:"valid_not_before"`
}

// FetchAndStoreCertificateRef fetches the CA certificate ref from NIOS WAPI and stores it
// as an environment variable. It makes a GET request to the cacertificate endpoint and
// extracts the _ref field from the first certificate in the response, then sets it as
// the environment variable specified by envVarName.
func FetchAndStoreCertificateRef(host, wapiVer, username, password, envVarName string) error {
	endpoint := fmt.Sprintf("%s/wapi/%s/cacertificate", host, wapiVer)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("fetchcertref: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // intentional for lab/CI grids with self-signed certs
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("fetchcertref: execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetchcertref: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fetchcertref: read response body: %w", err)
	}

	var certificates []CACertificate
	if err := json.Unmarshal(body, &certificates); err != nil {
		return fmt.Errorf("fetchcertref: unmarshal response: %w", err)
	}

	if len(certificates) == 0 {
		return fmt.Errorf("fetchcertref: no certificates found in response")
	}

	certRef := certificates[0].Ref
	if certRef == "" {
		return fmt.Errorf("fetchcertref: certificate ref is empty")
	}

	if err := os.Setenv(envVarName, certRef); err != nil {
		return fmt.Errorf("fetchcertref: failed to set environment variable %s: %w", envVarName, err)
	}

	return nil
}

type PreConfigClients struct {
	IPAM     *ipam.APIClient
	DHCP     *dhcp.APIClient
	DNS      *dns.APIClient
	SECURITY *security.APIClient
}

// PreConfig creates the network views required for integration testing.
// If a network view already exists (error contains "already exists"), it skips creation and continues.
// For any other error, it returns the error immediately.
func PreConfig(clients PreConfigClients) error {
	if clients.IPAM == nil {
		return fmt.Errorf("preconfig: IPAM client is required")
	}
	if clients.DHCP == nil {
		return fmt.Errorf("preconfig: DHCP client is required")
	}
	if clients.DNS == nil {
		return fmt.Errorf("preconfig: DNS client is required")
	}
	if clients.SECURITY == nil {
		return fmt.Errorf("preconfig: SECURITY client is required")
	}

	networkViews := []string{"test_network_view", "custom_view", "test_network_view2"}

	for _, viewName := range networkViews {
		networkViewBody := ipam.Networkview{
			Name: ipam.PtrString(viewName),
		}

		_, _, err := clients.IPAM.NetworkviewAPI.Create(context.Background()).
			Networkview(networkViewBody).
			Execute()

		if err != nil {
			// Check if the error is because the object already exists
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Network view %q already exists, skipping creation\n", viewName)
				continue
			}
			// Return any other error
			return fmt.Errorf("failed to create network view %q: %w", viewName, err)
		}

		fmt.Printf("Network view %q created successfully\n", viewName)
	}

	// Create networks
	type networkEntry struct {
		cidr        string
		networkView string
	}
	networks := []networkEntry{
		{"10.0.0.0/24", "default"},
		{"15.0.0.0/24", "default"},
		{"16.0.0.0/24", "default"},
	}

	for _, n := range networks {
		networkBody := ipam.Network{
			Network: &ipam.NetworkNetwork{
				String: ipam.PtrString(n.cidr),
			},
			NetworkView: ipam.PtrString(n.networkView),
			Comment:     ipam.PtrString("Created For Integration Testing"),
		}

		_, _, err := clients.IPAM.NetworkAPI.Create(context.Background()).
			Network(networkBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Network %q already exists, skipping creation\n", n.cidr)
				continue
			}
			return fmt.Errorf("failed to create network %q: %w", n.cidr, err)
		}

		fmt.Printf("Network %q created successfully\n", n.cidr)
	}

	// Create filter options
	filterOptions := []string{"example-option-filter-1", "example-option-filter-2"}

	for _, filterName := range filterOptions {
		filterOptionBody := dhcp.Filteroption{
			Name: dhcp.PtrString(filterName),
		}

		_, _, err := clients.DHCP.FilteroptionAPI.Create(context.Background()).
			Filteroption(filterOptionBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Filter option %q already exists, skipping creation\n", filterName)
				continue
			}
			return fmt.Errorf("failed to create filter option %q: %w", filterName, err)
		}

		fmt.Printf("Filter option %q created successfully\n", filterName)
	}

	// Create MAC filters
	macFilters := []string{"mac_filter", "mac_filter2", "example-mac-filter-1"}

	for _, macFilterName := range macFilters {
		filterMACBody := dhcp.Filtermac{
			Name: dhcp.PtrString(macFilterName),
		}

		_, _, err := clients.DHCP.FiltermacAPI.Create(context.Background()).
			Filtermac(filterMACBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("MAC filter %q already exists, skipping creation\n", macFilterName)
				continue
			}
			return fmt.Errorf("failed to create MAC filter %q: %w", macFilterName, err)
		}

		fmt.Printf("MAC filter %q created successfully\n", macFilterName)
	}

	// Create NAC filters
	nacFilters := []string{"nac_filter", "nac_filter_rule", "ipv6_nac_filter"}

	for _, nacFilterName := range nacFilters {
		filterNACBody := dhcp.Filternac{
			Name: dhcp.PtrString(nacFilterName),
		}

		_, _, err := clients.DHCP.FilternacAPI.Create(context.Background()).
			Filternac(filterNACBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("NAC filter %q already exists, skipping creation\n", nacFilterName)
				continue
			}
			return fmt.Errorf("failed to create NAC filter %q: %w", nacFilterName, err)
		}

		fmt.Printf("NAC filter %q created successfully\n", nacFilterName)
	}

	// Create IPv6 filter options
	ipv6FilterOptions := []string{"ipv6_option_filter", "ipv6_option_filter1"}

	for _, ipv6FilterName := range ipv6FilterOptions {
		ipv6FilterOptionBody := dhcp.Ipv6filteroption{
			Name: dhcp.PtrString(ipv6FilterName),
		}

		_, _, err := clients.DHCP.Ipv6filteroptionAPI.Create(context.Background()).
			Ipv6filteroption(ipv6FilterOptionBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("IPv6 filter option %q already exists, skipping creation\n", ipv6FilterName)
				continue
			}
			return fmt.Errorf("failed to create IPv6 filter option %q: %w", ipv6FilterName, err)
		}

		fmt.Printf("IPv6 filter option %q created successfully\n", ipv6FilterName)
	}

	// Create relay agent filters
	relayAgentFilters := []string{"relay_agent_filter", "relay_agent_filter2"}

	for _, relayFilterName := range relayAgentFilters {
		filterRelayAgentBody := dhcp.Filterrelayagent{
			Name: dhcp.PtrString(relayFilterName),
		}

		_, _, err := clients.DHCP.FilterrelayagentAPI.Create(context.Background()).
			Filterrelayagent(filterRelayAgentBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Relay agent filter %q already exists, skipping creation\n", relayFilterName)
				continue
			}
			return fmt.Errorf("failed to create relay agent filter %q: %w", relayFilterName, err)
		}

		fmt.Printf("Relay agent filter %q created successfully\n", relayFilterName)
	}

	// Create fingerprint filters
	fingerprintFilters := []string{"test_filter_fingerprint", "test_filter_fingerprint1"}

	for _, fpFilterName := range fingerprintFilters {
		filterFingerprintBody := dhcp.Filterfingerprint{
			Name: dhcp.PtrString(fpFilterName),
		}

		_, _, err := clients.DHCP.FilterfingerprintAPI.Create(context.Background()).
			Filterfingerprint(filterFingerprintBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Fingerprint filter %q already exists, skipping creation\n", fpFilterName)
				continue
			}
			return fmt.Errorf("failed to create fingerprint filter %q: %w", fpFilterName, err)
		}

		fmt.Printf("Fingerprint filter %q created successfully\n", fpFilterName)
	}

	// Create admin users
	adminUsers := []string{"aws1", "aws2"}

	for _, adminUserName := range adminUsers {
		adminUser := security.Adminuser{
			Name:        security.PtrString(adminUserName),
			AuthType:    security.PtrString("SAML"),
			AdminGroups: []string{"cloud-api-only"},
		}

		_, _, err := clients.SECURITY.AdminuserAPI.Create(context.Background()).
			Adminuser(adminUser).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Admin user %q already exists, skipping creation\n", adminUserName)
				continue
			}
			return fmt.Errorf("failed to create admin user %q: %w", adminUserName, err)
		}

		fmt.Printf("Admin user %q created successfully\n", adminUserName)
	}

	// Create DNS views
	dnsViews := []string{"custom_dns_view"}

	for _, dnsViewName := range dnsViews {
		dnsViewBody := dns.View{
			Name: dns.PtrString(dnsViewName),
		}

		_, _, err := clients.DNS.ViewAPI.Create(context.Background()).
			View(dnsViewBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DNS view %q already exists, skipping creation\n", dnsViewName)
				continue
			}
			return fmt.Errorf("failed to create DNS view %q: %w", dnsViewName, err)
		}

		fmt.Printf("DNS view %q created successfully\n", dnsViewName)
	}

	// Create DDNS principal cluster groups
	ddnsPrincipalClusterGroups := []string{"dynamic_update_grp_1", "dynamic_update_grp_2"}

	for _, groupName := range ddnsPrincipalClusterGroups {
		ddnsPrincipalClusterGroupBody := dns.DdnsPrincipalclusterGroup{
			Name: dns.PtrString(groupName),
		}

		_, _, err := clients.DNS.DdnsPrincipalclusterGroupAPI.Create(context.Background()).
			DdnsPrincipalclusterGroup(ddnsPrincipalClusterGroupBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DDNS principal cluster group %q already exists, skipping creation\n", groupName)
				continue
			}
			return fmt.Errorf("failed to create DDNS principal cluster group %q: %w", groupName, err)
		}

		fmt.Printf("DDNS principal cluster group %q created successfully\n", groupName)
	}

	// Create DNS64 groups
	dns64Groups := []struct {
		name   string
		prefix string
	}{
		{name: "dns64_group", prefix: "64:FF9B::/96"},
	}

	for _, g := range dns64Groups {
		dns64GroupBody := dns.Dns64group{
			Name:   dns.PtrString(g.name),
			Prefix: dns.PtrString(g.prefix),
		}

		_, _, err := clients.DNS.Dns64groupAPI.Create(context.Background()).
			Dns64group(dns64GroupBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DNS64 group %q already exists, skipping creation\n", g.name)
				continue
			}
			return fmt.Errorf("failed to create DNS64 group %q: %w", g.name, err)
		}

		fmt.Printf("DNS64 group %q created successfully\n", g.name)
	}

	// Create stub NS groups
	stubNSGroups := []string{"stub_ns_group1", "stub_ns_group2"}

	for _, groupName := range stubNSGroups {
		stubMemberBody := dns.NsgroupStubmember{
			Name: dns.PtrString(groupName),
			StubMembers: []dns.NsgroupStubmemberStubMembers{
				{
					Name: dns.PtrString("infoblox.member"),
				},
			},
		}

		_, _, err := clients.DNS.NsgroupStubmemberAPI.Create(context.Background()).
			NsgroupStubmember(stubMemberBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Stub NS group %q already exists, skipping creation\n", groupName)
				continue
			}
			return fmt.Errorf("failed to create stub NS group %q: %w", groupName, err)
		}

		fmt.Printf("Stub NS group %q created successfully\n", groupName)
	}

	// Create forward stub servers
	forwardStubServers := []string{"ensg1", "ensg2"}

	for _, serverName := range forwardStubServers {
		forwardStubServerBody := dns.NsgroupForwardstubserver{
			Name: dns.PtrString(serverName),
		}

		_, _, err := clients.DNS.NsgroupForwardstubserverAPI.Create(context.Background()).
			NsgroupForwardstubserver(forwardStubServerBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Forward stub server %q already exists, skipping creation\n", serverName)
				continue
			}
			return fmt.Errorf("failed to create forward stub server %q: %w", serverName, err)
		}

		fmt.Printf("Forward stub server %q created successfully\n", serverName)
	}

	// Create IPv4 reverse mapping zone auth
	zoneAuthBody := dns.ZoneAuth{
		Fqdn:       dns.PtrString("192.168.10.0/24"),
		View:       dns.PtrString("default"),
		ZoneFormat: dns.PtrString("IPV4"),
	}

	_, _, err := clients.DNS.ZoneAuthAPI.Create(context.Background()).
		ZoneAuth(zoneAuthBody).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Printf("Zone auth %q already exists, skipping creation\n", "192.168.10.0/24")
		} else {
			return fmt.Errorf("failed to create zone auth %q: %w", "192.168.10.0/24", err)
		}
	} else {
		fmt.Printf("Zone auth %q created successfully\n", "192.168.10.0/24")
	}

	// Create IPv6 reverse mapping zone auth
	ipv6ZoneAuthBody := dns.ZoneAuth{
		Fqdn:       dns.PtrString("2001::/64"),
		View:       dns.PtrString("default"),
		ZoneFormat: dns.PtrString("IPV6"),
	}

	_, _, err = clients.DNS.ZoneAuthAPI.Create(context.Background()).
		ZoneAuth(ipv6ZoneAuthBody).
		Execute()

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Printf("Zone auth %q already exists, skipping creation\n", "2001::/64")
		} else {
			return fmt.Errorf("failed to create zone auth %q: %w", "2001::/64", err)
		}
	} else {
		fmt.Printf("Zone auth %q created successfully\n", "2001::/64")
	}

	// Create DHCP failovers
	failovers := []struct {
		name      string
		secondary string
	}{
		{name: "example_failover_association", secondary: "2.2.2.2"},
		{name: "example_failover_association1", secondary: "2.2.2.3"},
	}

	for _, f := range failovers {
		failoverBody := dhcp.Dhcpfailover{
			Name:                dhcp.PtrString(f.name),
			PrimaryServerType:   dhcp.PtrString("GRID"),
			Primary:             dhcp.PtrString("infoblox.member"),
			SecondaryServerType: dhcp.PtrString("EXTERNAL"),
			Secondary:           dhcp.PtrString(f.secondary),
		}

		_, _, err := clients.DHCP.DhcpfailoverAPI.Create(context.Background()).
			Dhcpfailover(failoverBody).
			Execute()

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				fmt.Printf("DHCP failover %q already exists, skipping creation\n", f.name)
				continue
			}
			return fmt.Errorf("failed to create DHCP failover %q: %w", f.name, err)
		}

		fmt.Printf("DHCP failover %q created successfully\n", f.name)
	}

	return nil
}

func main() {
	host := ""
	wapiVer := ""
	username := ""
	password := ""

	certUpload, err := UploadInit(host, wapiVer, username, password)
	if err != nil {
		fmt.Printf("Error initializing upload: %v\n", err)
		return
	}
	fmt.Printf("Upload initialized: %+v\n", certUpload)

	err = UploadFile(host, wapiVer, username, password, certUpload.URL, "./nios_security_certificate_authservice/cert.pem")
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
	fmt.Println("File uploaded successfully")

	err = UploadCertificate(host, wapiVer, username, password, "EAP_CA", "infoblox.localdomain", certUpload.Token)
	if err != nil {
		fmt.Printf("Error uploading certificate: %v\n", err)
		return
	}
	fmt.Println("Certificate uploaded successfully")

	err = FetchAndStoreCertificateRef(host, wapiVer, username, password, "NIOS_CA_CERT1_REF")
	if err != nil {
		fmt.Printf("Error fetching and storing certificate ref: %v\n", err)
		return
	}
	fmt.Println("Certificate ref fetched and stored successfully")
}
