// client.go
package smsgw

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// Client represents a Puzzel SMS Gateway client
type Client struct {
	baseURL        string
	serviceID      int
	username       string
	password       string
	batchReference string
	httpClient     *http.Client
}

// WithBatchReference sets a custom batch reference for all requests
func WithBatchReference(batchReference string) ClientOption {
	return func(c *Client) {
		c.batchReference = batchReference
	}
}

// WithTimeout sets a custom timeout for the HTTP client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient initializes and returns a new SMS Gateway client
func NewClient(baseURL string, serviceID int, username, password string, options ...ClientOption) *Client {
	client := &Client{
		baseURL:   baseURL,
		serviceID: serviceID,
		username:  username,
		password:  password,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	
	// Apply options
	for _, option := range options {
		option(client)
	}
	
	return client
}

// Send is a simplified method to send one or more messages
func (c *Client) Send(ctx context.Context, messages []Message) (*SmsGatewayResponse, error) {
	return c.SendMessages(ctx, messages, c.batchReference)
}

// SendMessages sends messages to the SMS gateway with an explicit batch reference
func (c *Client) SendMessages(ctx context.Context, messages []Message, batchReference string) (*SmsGatewayResponse, error) {
	requestBody := map[string]interface{}{
		"serviceId":      c.serviceID,
		"username":       c.username,
		"password":       c.password,
		"batchReference": batchReference,
		"message":        messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Failed to marshal request body: %v", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/gw/rs/sendMessages", bytes.NewReader(jsonData))
	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	log.Printf("Sending request to %s", req.URL)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s", resp.Status)
		return nil, &APIError{StatusCode: resp.StatusCode, Message: resp.Status}
	}

	var gatewayResponse SmsGatewayResponse
	if err := json.NewDecoder(resp.Body).Decode(&gatewayResponse); err != nil {
		log.Printf("Failed to decode response body: %v", err)
		return nil, err
	}

	log.Printf("Received response: %+v", gatewayResponse)
	return &gatewayResponse, nil
}
