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
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
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
	IPAM *ipam.APIClient
	DHCP *dhcp.APIClient
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

	return nil
}

func main() {
	host := "https://172.28.83.140"
	wapiVer := "v2.13.6"
	username := "admin"
	password := "Infoblox@123"

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
