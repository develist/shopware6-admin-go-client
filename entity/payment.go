package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// PaymentMethod data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type PaymentMethod struct {
	Active                         bool              `json:"active,omitempty"`
	AfterOrderEnabled              bool              `json:"afterOrderEnabled,omitempty"`
	AppPaymentMethod               *AppPaymentMethod `json:"appPaymentMethod,omitempty"`
	AvailabilityRule               *Rule             `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             string            `json:"availabilityRuleId,omitempty"`
	CreatedAt                      *time.Time        `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField    `json:"customFields,omitempty"`
	Customers                      *Customer         `json:"customers,omitempty"`
	Description                    string            `json:"description,omitempty"`
	DistinguishableName            string            `json:"distinguishableName,omitempty"`
	FormattedHandlerIdentifier     string            `json:"formattedHandlerIdentifier,omitempty"`
	Id                             string            `json:"id,omitempty"`
	Media                          *Media            `json:"media,omitempty"`
	MediaId                        string            `json:"mediaId,omitempty"`
	Name                           string            `json:"name,omitempty"`
	OrderTransactions              *OrderTransaction `json:"orderTransactions,omitempty"`
	Plugin                         *Plugin           `json:"plugin,omitempty"`
	PluginId                       string            `json:"pluginId,omitempty"`
	Position                       int               `json:"position,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel     `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  *SalesChannel     `json:"salesChannels,omitempty"`
	Translated                     *interface{}      `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                      *time.Time        `json:"updatedAt,omitempty"`
}
