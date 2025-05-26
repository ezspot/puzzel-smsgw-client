// examples/send_sms.go
package main

import (
	"context"
	"log"
	"time"

	smsgw "github.com/ezspot/puzzel-smsgw-client"
)

func main() {
	// Example 1: Basic usage with default client
	basicExample()

	// Example 2: Advanced usage with all options
	advancedExample()
}

func basicExample() {
	// Initialize client with default settings
	client := smsgw.NewClient(
		"https://api.puzzel.com", 
		1000, 
		"username", 
		"password",
		smsgw.WithBatchReference("basic-example-batch"),
	)

	// Create a simple message
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

	// Send using the simplified Send method
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Send(ctx, messages)
	if err != nil {
		log.Fatalf("Error sending messages: %v", err)
	}

	log.Printf("Basic example response: %+v", response)
}

func advancedExample() {
	// Initialize client with custom options
	client := smsgw.NewClient(
		"https://api.puzzel.com",
		1000,
		"username",
		"password",
		smsgw.WithTimeout(15*time.Second),
		smsgw.WithBatchReference("advanced-example-batch"),
	)

	// Create a message with advanced settings
	messages := []smsgw.Message{
		{
			Recipient:       "+4712345678",
			Content:         "Advanced message example!",
			Price:           100, // Optional price
			ClientReference: "client-ref-123", // Optional client reference
			Settings: &smsgw.Settings{
				Priority:           1,
				Validity:           173,
				Differentiator:     "sms group 1",
				InvoiceNode:        "marketing department",
				Age:                18,
				NewSession:         true,
				SessionID:          "01bxmt7f8b8h3zkwe2vg",
				AutoDetectEncoding: true,
				OriginatorSettings: &smsgw.OriginatorSettings{
					Originator:     "1960",
					OriginatorType: "NETWORK",
				},
				GasSettings: &smsgw.GasSettings{
					ServiceCode: "02001",
					Description: "SMS",
				},
				SendWindow: &smsgw.SendWindow{
					StartDate: "2025-05-27",
					StopDate:  "2025-05-27",
					StartTime: "10:00:00",
					StopTime:  "18:00:00",
				},
				Parameter: &smsgw.Parameter{
					BusinessModel:            "contact center",
					Flash:                     true,
					ParsingType:               "AUTO_DETECT",
					SkipCustomerReportDelivery: true,
				},
			},
		},
	}

	// Send using the explicit SendMessages method with a custom batch reference
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.SendMessages(ctx, messages, "custom-batch-ref-123")
	if err != nil {
		log.Fatalf("Error sending messages: %v", err)
	}

	log.Printf("Advanced example response: %+v", response)
}
