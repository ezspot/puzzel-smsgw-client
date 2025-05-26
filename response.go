// response.go
package smsgw

type SmsGatewayResponse struct {
	BatchReference string          `json:"batchReference"`
	MessageStatus  []MessageStatus `json:"messageStatus"`
}

type MessageStatus struct {
	StatusCode      int    `json:"statusCode"`
	StatusMessage   string `json:"statusMessage"`
	ClientReference string `json:"clientReference"`
	Recipient       string `json:"recipient"`
	MessageID       string `json:"messageId"`
	SessionID       string `json:"sessionId"`
	SequenceIndex   int    `json:"sequenceIndex"`
}
