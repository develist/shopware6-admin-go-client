package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Webhook data structure
// Added since version: 6.3.1.0
// Required fields: name, eventName, url, errorCount, createdAt
type Webhook struct {
	Active     *bool      `json:"active,omitempty"`
	App        *App       `json:"app,omitempty"`
	AppId      *string    `json:"appId,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	ErrorCount int        `json:"errorCount,omitempty"`
	EventName  string     `json:"eventName,omitempty"`
	Id         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	Url        string     `json:"url,omitempty"`
}

// WebhookEventLog data structure
// Added since version: 6.4.1.0
// Required fields: webhookName, eventName, deliveryStatus, url, createdAt
type WebhookEventLog struct {
	AppName              *string        `json:"appName,omitempty"`
	AppVersion           *string        `json:"appVersion,omitempty"`
	CreatedAt            *time.Time     `json:"createdAt,omitempty"`
	CustomFields         *[]CustomField `json:"customFields,omitempty"`
	DeliveryStatus       string         `json:"deliveryStatus,omitempty"`
	EventName            string         `json:"eventName,omitempty"`
	Id                   string         `json:"id,omitempty"`
	ProcessingTime       *int           `json:"processingTime,omitempty"`
	RequestContent       *interface{}   `json:"requestContent,omitempty"`  // map[type:object]
	ResponseContent      *interface{}   `json:"responseContent,omitempty"` // map[type:object]
	ResponseReasonPhrase *string        `json:"responseReasonPhrase,omitempty"`
	ResponseStatusCode   *int           `json:"responseStatusCode,omitempty"`
	Timestamp            *int           `json:"timestamp,omitempty"`
	UpdatedAt            *time.Time     `json:"updatedAt,omitempty"`
	Url                  string         `json:"url,omitempty"`
	WebhookName          string         `json:"webhookName,omitempty"`
}
