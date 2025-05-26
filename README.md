# Puzzel SMS Gateway Go Client

A Go client library for interacting with the Puzzel SMS Gateway API. This library provides a simple and efficient way to send SMS messages through Puzzel's messaging service.

## Features

- Send single or multiple SMS messages in one batch
- Comprehensive message configuration options (originator, type, scheduling, etc.)
- Flexible client options with functional options pattern
- Support for Strex and advanced SMS features
- Scheduled message delivery with time windows
- Context support for request cancellation and timeouts
- Error handling for API responses
- Thread-safe client implementation

## Installation

To use this library in your Go project, run:

```bash
go get github.com/ezspot/puzzel-smsgw-client
```

## Usage

### Basic Example

```go
package main

import (
	"context"
	"log"
	"time"

	smsgw "github.com/ezspot/puzzel-smsgw-client"
)

func main() {
	// Initialize client with your Puzzel credentials
	client := smsgw.NewClient(
		"https://api.puzzel.com", // Puzzel API base URL
		1000,                     // Service ID
		"your-username",          // Your Puzzel username
		"your-password",          // Your Puzzel password
		// Optional client options
		smsgw.WithBatchReference("my-batch-reference"),
		smsgw.WithTimeout(15*time.Second),
	)

	// Prepare your message(s)
	messages := []smsgw.Message{
		{
			Recipient: "+4712345678",
			Content:   "Hello from Puzzel!",
			Settings: &smsgw.Settings{
				OriginatorSettings: &smsgw.OriginatorSettings{
					Originator:     "Puzzel",
					OriginatorType: "ALPHANUMERIC",
				},
			},
		},
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send messages using the simplified Send method
	response, err := client.Send(ctx, messages)
	if err != nil {
		log.Fatalf("Error sending messages: %v", err)
	}
	
	// Or send with an explicit batch reference
	// response, err := client.SendMessages(ctx, messages, "batch-123")

	log.Printf("Send response: %+v", response)
}
```

## Message Structure

The `Message` struct has the following fields:

- `Recipient` (string): The phone number of the message recipient (in international format)
- `Content` (string): The message content
- `Price` (int): Optional price for the message
- `ClientReference` (string): Optional client-side reference ID
- `Settings` (*Settings): Optional message settings

### Settings

The `Settings` struct includes:

- `Priority` (int): Message priority level
- `Validity` (int): Message validity period
- `Differentiator` (string): Custom differentiator for message grouping
- `Age` (int): Age restriction for content
- `NewSession` (bool): Whether to create a new session
- `SessionID` (string): ID for an existing session
- `InvoiceNode` (string): Invoice node for billing
- `AutoDetectEncoding` (bool): Whether to auto-detect message encoding
- `OriginatorSettings`: Configuration for the message sender
  - `Originator`: The sender ID or phone number
  - `OriginatorType`: Type of originator (e.g., "ALPHANUMERIC", "NUMERIC")
- `GasSettings`: Configuration for GAS (Gateway Application Services)
  - `ServiceCode`: The service code
  - `Description`: Optional service description
- `SendWindow`: Configuration for scheduled message delivery
  - `StartDate`: Start date in YYYY-MM-DD format
  - `StopDate`: Optional end date in YYYY-MM-DD format
  - `StartTime`: Start time in HH:MM:SS format
  - `StopTime`: Optional end time in HH:MM:SS format
- `Parameter`: Additional configuration parameters
  - `BusinessModel`: Business model identifier
  - `Dcs`: Data coding scheme
  - `Udh`: User data header
  - `Pid`: Protocol identifier
  - `Flash`: Whether message is a flash SMS
  - `ParsingType`: Content parsing type
  - `SkipCustomerReportDelivery`: Whether to skip delivery reports
  - Various Strex-related parameters for payment services

## Error Handling

The library returns standard Go errors for network and API-related issues. The `SendMessages` function returns a `SmsGatewayResponse` that includes detailed information about the operation's result.

## Client Options

The client supports functional options for flexible configuration:

```go
// Create a client with custom options
client := smsgw.NewClient(
    "https://api.puzzel.com",
    1000,
    "username",
    "password",
    // Optional configurations
    smsgw.WithBatchReference("my-default-batch"),
    smsgw.WithTimeout(15*time.Second),
)
```

Available options:

- `WithBatchReference(ref string)`: Set a default batch reference for all requests
- `WithTimeout(duration time.Duration)`: Set a custom HTTP client timeout
- `WithHTTPClient(client *http.Client)`: Use a custom HTTP client

## Advanced Usage Examples

See the `example/send_sms.go` file for comprehensive examples including:

- Basic message sending
- Advanced message configuration
- Custom client options
- Scheduled message delivery
- Parameter configuration

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, please contact Puzzel support or open an issue in the GitHub repository.
