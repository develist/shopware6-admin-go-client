package entity

import "time"

// completed

type Webhook struct {
	Active     bool      `json:"active,omitempty"`
	App        *App      `json:"app,omitempty"`
	AppId      string    `json:"appId,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	ErrorCount int       `json:"errorCount,omitempty"`
	EventName  string    `json:"eventName,omitempty"`
	Id         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	Url        string    `json:"url,omitempty"`
}
type WebhookEventLog struct {
	AppName        string         `json:"appName,omitempty"`
	AppVersion     string         `json:"appVersion,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	DeliveryStatus string         `json:"deliveryStatus,omitempty"`
	EventName      string         `json:"eventName,omitempty"`
	Id             string         `json:"id,omitempty"`
	ProcessingTime int            `json:"processingTime,omitempty"`
	// RequestContent map[type:object] `json:"requestContent,omitempty"`
	// ResponseContent map[type:object] `json:"responseContent,omitempty"`
	ResponseReasonPhrase string    `json:"responseReasonPhrase,omitempty"`
	ResponseStatusCode   int       `json:"responseStatusCode,omitempty"`
	Timestamp            int       `json:"timestamp,omitempty"`
	UpdatedAt            time.Time `json:"updatedAt,omitempty"`
	Url                  string    `json:"url,omitempty"`
	WebhookName          string    `json:"webhookName,omitempty"`
}
