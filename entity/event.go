package entity

import "time"

// completed

type EventAction struct {
	ActionName string `json:"actionName,omitempty"`
	Active     bool   `json:"active,omitempty"`
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt     *time.Time     `json:"createdAt,omitempty"`
	CustomFields  *[]CustomField `json:"customFields,omitempty"`
	EventName     string         `json:"eventName,omitempty"`
	Id            string         `json:"id,omitempty"`
	Rules         *Rule          `json:"rules,omitempty"`
	SalesChannels *SalesChannel  `json:"salesChannels,omitempty"`
	Title         string         `json:"title,omitempty"`
	UpdatedAt     *time.Time     `json:"updatedAt,omitempty"`
}
type EventActionRule struct {
	EventAction   *EventAction `json:"eventAction,omitempty"`
	EventActionId string       `json:"eventActionId,omitempty"`
	Id            string       `json:"id,omitempty"`
	Rule          *Rule        `json:"rule,omitempty"`
	RuleId        string       `json:"ruleId,omitempty"`
}
type EventActionSalesChannel struct {
	EventAction    *EventAction  `json:"eventAction,omitempty"`
	EventActionId  string        `json:"eventActionId,omitempty"`
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}
