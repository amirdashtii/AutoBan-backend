package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/amirdashtii/AutoBan/pkg/logger"
)

// HTTPClient interface for making HTTP requests
type HTTPClient interface {
	Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error)
	Put(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error)
	Delete(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
}

// Client implements HTTPClient interface
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new HTTP client
func NewClient(baseURL string, timeout time.Duration) HTTPClient {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// Get makes a GET request
func (c *Client) Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
	return c.makeRequest(ctx, "GET", url, nil, headers)
}

// Post makes a POST request
func (c *Client) Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return c.makeRequest(ctx, "POST", url, body, headers)
}

// Put makes a PUT request
func (c *Client) Put(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return c.makeRequest(ctx, "PUT", url, body, headers)
}

// Delete makes a DELETE request
func (c *Client) Delete(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
	return c.makeRequest(ctx, "DELETE", url, nil, headers)
}

// makeRequest is the internal method that handles all HTTP requests
func (c *Client) makeRequest(ctx context.Context, method, url string, body interface{}, headers map[string]string) (*http.Response, error) {
	var reqBody io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			logger.Error(err, "Failed to marshal request body")
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, reqBody)
	if err != nil {
		logger.Error(err, "Failed to create HTTP request")
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Error(err, "Failed to make HTTP request")
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	return resp, nil
}

// ParseResponse parses HTTP response body into a struct
func ParseResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err, "Failed to read response body")
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		logger.Error(err, "Failed to unmarshal response body")
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
