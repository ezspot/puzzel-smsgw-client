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

type Client struct {
	baseURL    string
	serviceID  int
	username   string
	password   string
	httpClient *http.Client
}

// NewClient initializes and returns a new SMS Gateway client
func NewClient(baseURL string, serviceID int, username, password string) *Client {
	return &Client{
		baseURL:   baseURL,
		serviceID: serviceID,
		username:  username,
		password:  password,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

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
