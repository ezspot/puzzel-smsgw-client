// message.go
package smsgw

import "time"

type Message struct {
	Recipient       string    `json:"recipient"`
	Content         string    `json:"content"`
	Price           int       `json:"price,omitempty"`
	ClientReference string    `json:"clientReference,omitempty"`
	Settings        *Settings `json:"settings,omitempty"`
}

type Settings struct {
	Priority                   int                 `json:"priority,omitempty"`
	Validity                   int                 `json:"validity,omitempty"`
	Differentiator             string              `json:"differentiator,omitempty"`
	Age                        int                 `json:"age,omitempty"`
	NewSession                 bool                `json:"newSession,omitempty"`
	SessionID                  string              `json:"sessionId,omitempty"`
	InvoiceNode                string              `json:"invoiceNode,omitempty"`
	AutoDetectEncoding         bool                `json:"autoDetectEncoding,omitempty"`
	SafeRemoveNonGsmCharacters bool                `json:"safeRemoveNonGsmCharacters,omitempty"`
	OriginatorSettings         *OriginatorSettings `json:"originatorSettings,omitempty"`
	GasSettings                *GasSettings        `json:"gasSettings,omitempty"`
	SendWindow                 *SendWindow         `json:"sendWindow,omitempty"`
	Parameters                 map[string]string   `json:"parameter,omitempty"`
}

type OriginatorSettings struct {
	OriginatorType string `json:"originatorType"`
	Originator     string `json:"originator"`
}

type GasSettings struct {
	ServiceCode string `json:"serviceCode"`
	Description string `json:"description,omitempty"`
}

type SendWindow struct {
	StartDate time.Time  `json:"startDate"`
	StartTime *time.Time `json:"startTime,omitempty"`
	StopDate  *time.Time `json:"stopDate,omitempty"`
	StopTime  *time.Time `json:"stopTime,omitempty"`
}
