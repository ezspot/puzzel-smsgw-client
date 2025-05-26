// message.go
package smsgw

// Package smsgw provides a client for the Puzzel SMS Gateway API

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
	Parameter                  *Parameter          `json:"parameter,omitempty"`
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
	StartDate string `json:"startDate"`
	StopDate  string `json:"stopDate,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	StopTime  string `json:"stopTime,omitempty"`
}

type Parameter struct {
	BusinessModel                string `json:"businessModel,omitempty"`
	Dcs                         string `json:"dcs,omitempty"`
	Udh                         string `json:"udh,omitempty"`
	Pid                         int    `json:"pid,omitempty"`
	Flash                       bool   `json:"flash,omitempty"`
	ParsingType                 string `json:"parsingType,omitempty"`
	SkipCustomerReportDelivery  bool   `json:"skipCustomerReportDelivery,omitempty"`
	StrexVerificationTimeout    string `json:"strexVerificationTimeout,omitempty"`
	StrexMerchantSellOption     string `json:"strexMerchantSellOption,omitempty"`
	StrexConfirmChannel         string `json:"strexConfirmChannel,omitempty"`
	StrexAuthorizationToken     string `json:"strexAuthorizationToken,omitempty"`
}
