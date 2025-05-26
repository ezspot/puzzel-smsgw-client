// examples/send_sms.go
package main

import (
	"context"
	"log"
	"time"

	"puzzel-smsgw/smsgw"
)

func main() {
	client := smsgw.NewClient("https://api.puzzel.com", 1000, "username", "password")

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.SendMessages(ctx, messages, "batch-123")
	if err != nil {
		log.Fatalf("Error sending messages: %v", err)
	}

	log.Printf("Send response: %+v", response)
}
