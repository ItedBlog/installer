package units

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"os"
)

// GitHubReleaseResponse 定义 GitHub API 返回的版本信息结构
// GitHubReleaseResponse defines the structure of GitHub API release response
type GitHubReleaseResponse struct {
	ID string `json:"id"` // 版本标签，如 "v1.0.0"
	Name    string `json:"name"`    // 发布名称
	Body    string `json:"body"`    // 发布说明
}

// getVersion fetches the latest version number from GitHub repository
// Returns: version string, e.g. "v1.0.0"
func getVersion() string {
	// Set HTTP request headers for GitHub API v3 JSON response
	headers := http.Header{
		"Accept":               []string{"application/vnd.github.v3+json"},
		"x-github-api-version": []string{"2026-03-10"},
	}
	
	// Create HTTP client with timeout (5 seconds)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	
	// Create GET request to GitHub API mirror
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/ItedBlog/installer/releases/latest", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create HTTP request: %v\n", err)
		os.Exit(1)
	}
	
	// Set request headers
	req.Header = headers
	
	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send HTTP request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	
	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Failed to get GitHub API response: %s\n", resp.Status)
		os.Exit(1)
	}
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read HTTP response body: %v\n", err)
		os.Exit(1)
	}
	
	// Parse JSON response
	var release GitHubReleaseResponse
	if err := json.Unmarshal(body, &release); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse GitHub API response JSON: %v\n", err)
		os.Exit(1)
	}
	

	// Return version tag
	if release.ID == "" {
		return nil
	}
	return release.ID
}

func getLatestPackageURL() string {
	// Set HTTP request headers for GitHub API v3 JSON response
	headers := http.Header{
		"Accept":               []string{"application/vnd.github.v3+json"},
		"x-github-api-version": []string{"2026-03-10"},
	}
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	// Create GET request to GitHub API mirror
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/ItedBlog/installer/releases/latest", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create HTTP request: %v\n", err)
		os.Exit(1)
	}
	
	// Set request headers
	req.Header = headers
	
	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send HTTP request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	
	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Failed to get GitHub API response: %s\n", resp.Status)
		os.Exit(1)
	}
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read HTTP response body: %v\n", err)
		os.Exit(1)
	}
	
	// Parse JSON response
	var release GitHubReleaseResponse
	if err := json.Unmarshal(body, &release); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse GitHub API response JSON: %v\n", err)
		os.Exit(1)
	}
	
	// Return latest package URL
	if release.Body == "" {
		return nil
	}
	return release.Body
}