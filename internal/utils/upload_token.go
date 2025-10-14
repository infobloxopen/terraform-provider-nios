package utils

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

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UploadInitResponse struct represents the response from the NIOS uploadinit API call
type UploadInitResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// GenerateUploadToken generates an upload token and URL by calling the NIOS uploadinit API
func GenerateUploadToken(ctx context.Context, baseURL, username, password string) (*UploadInitResponse, error) {
	var uploadInitResponse UploadInitResponse
	// Extract configuration values using reflection for flexibility
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Generate upload token and URL by calling uploadinit
	uploadInitURL := fmt.Sprintf("%s/wapi/v2.13.6/fileop?_function=uploadinit", baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", uploadInitURL, bytes.NewReader([]byte("{}")))
	if err != nil {
		return &uploadInitResponse, fmt.Errorf("error creating uploadinit request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	tflog.Debug(ctx, fmt.Sprintf("Making uploadinit request to: %s", uploadInitURL))
	resp, err := httpClient.Do(req)
	if err != nil {
		return &uploadInitResponse, fmt.Errorf("error making uploadinit request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &uploadInitResponse, fmt.Errorf("uploadinit request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(&uploadInitResponse); err != nil {
		return &uploadInitResponse, fmt.Errorf("error decoding uploadinit response: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Generated upload token: %s with URL: %s", uploadInitResponse.Token, uploadInitResponse.URL))
	return &uploadInitResponse, nil
}

// UploadFile uploads a file to the Infoblox NIOS server using the provided upload URL.
func UploadFile(ctx context.Context, uploadURL, filePath, username, password string) error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}
	defer func() { _ = file.Close() }()
	// Create a buffer for the multipart form
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create the form file field
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("error creating form file: %w", err)
	}
	// Copy the file content to the form field
	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("error copying file content: %w", err)
	}
	//Close the multipart writer to finalize the form
	if err = writer.Close(); err != nil {
		return fmt.Errorf("error finalizing multipart form: %w", err)
	}

	// Create a new request to upload the file
	uploadReq, err := http.NewRequestWithContext(ctx, "POST", uploadURL, &requestBody)
	if err != nil {
		return fmt.Errorf("error creating file upload request: %w", err)
	}

	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.SetBasicAuth(username, password)

	tflog.Debug(ctx, fmt.Sprintf("Uploading the file to: %s", uploadURL))
	uploadResp, err := httpClient.Do(uploadReq)
	if err != nil {
		return fmt.Errorf("error uploading the file: %w", err)
	}
	defer func() { _ = uploadResp.Body.Close() }()
	if uploadResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(uploadResp.Body)
		return fmt.Errorf("file upload failed with status %d: %s", uploadResp.StatusCode, string(bodyBytes))
	}

	tflog.Info(ctx, fmt.Sprintf("file %s uploaded successfully", filePath))
	return nil
}

// UploadFileWithToken handles the complete process of generating an upload token and uploading a file
// It returns the token from the successful upload or an error if any step fails
func UploadFileWithToken(ctx context.Context, baseUrl, filePath, username, password string) (string, error) {
	// Generate the upload token
	uploadInitResponse, err := GenerateUploadToken(ctx, baseUrl, username, password)
	if err != nil {
		return "", fmt.Errorf("unable to generate upload token: %w", err)
	}

	// Upload the file using the token URL
	if err = UploadFile(ctx, uploadInitResponse.URL, filePath, username, password); err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return uploadInitResponse.Token, nil
}
