package entity

import "time"

// completed

type PaymentMethod struct {
	Active                         bool              `json:"active,omitempty"`
	AfterOrderEnabled              bool              `json:"afterOrderEnabled,omitempty"`
	AppPaymentMethod               *AppPaymentMethod `json:"appPaymentMethod,omitempty"`
	AvailabilityRule               *Rule             `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             string            `json:"availabilityRuleId,omitempty"`
	CreatedAt                      time.Time         `json:"createdAt,omitempty"`
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
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
