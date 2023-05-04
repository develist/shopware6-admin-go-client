package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// PaymentMethod data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type PaymentMethod struct {
	Active                         *bool             `json:"active,omitempty"`
	AfterOrderEnabled              *bool             `json:"afterOrderEnabled,omitempty"`
	AppPaymentMethod               *AppPaymentMethod `json:"appPaymentMethod,omitempty"`
	Asynchronous                   *bool             `json:"asynchronous,omitempty"`
	AvailabilityRule               *Rule             `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             *string           `json:"availabilityRuleId,omitempty"`
	CreatedAt                      *time.Time        `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField    `json:"customFields,omitempty"`
	Customers                      *Customer         `json:"customers,omitempty"`
	Description                    *string           `json:"description,omitempty"`
	DistinguishableName            *string           `json:"distinguishableName,omitempty"`
	FormattedHandlerIdentifier     *string           `json:"formattedHandlerIdentifier,omitempty"`
	HandlerIdentifier              *string           `json:"handlerIdentifier,omitempty"`
	Id                             string            `json:"id,omitempty"`
	Media                          *Media            `json:"media,omitempty"`
	MediaId                        *string           `json:"mediaId,omitempty"`
	Name                           string            `json:"name,omitempty"`
	OrderTransactions              *OrderTransaction `json:"orderTransactions,omitempty"`
	Plugin                         *Plugin           `json:"plugin,omitempty"`
	PluginId                       *string           `json:"pluginId,omitempty"`
	Position                       *int              `json:"position,omitempty"`
	Prepared                       *bool             `json:"prepared,omitempty"`
	Refundable                     *bool             `json:"refundable,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel     `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  *SalesChannel     `json:"salesChannels,omitempty"`
	Synchronous                    *bool             `json:"synchronous,omitempty"`
	Translated                     *interface{}      `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                      *time.Time        `json:"updatedAt,omitempty"`
}
