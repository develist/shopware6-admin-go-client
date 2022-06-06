package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Plugin data structure
// Added since version: 6.0.0.0
// Required fields: baseClass, name, autoload, version, createdAt, label
type Plugin struct {
	Active            bool           `json:"active,omitempty"`
	Author            string         `json:"author,omitempty"`
	Autoload          *interface{}   `json:"autoload,omitempty"` // map[type:object]
	BaseClass         string         `json:"baseClass,omitempty"`
	Changelog         *interface{}   `json:"changelog,omitempty"` // map[type:object]
	ComposerName      string         `json:"composerName,omitempty"`
	Copyright         string         `json:"copyright,omitempty"`
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Description       string         `json:"description,omitempty"`
	Icon              string         `json:"icon,omitempty"`
	Id                string         `json:"id,omitempty"`
	InstalledAt       *time.Time     `json:"installedAt,omitempty"`
	Label             string         `json:"label,omitempty"`
	License           string         `json:"license,omitempty"`
	ManagedByComposer bool           `json:"managedByComposer,omitempty"`
	ManufacturerLink  string         `json:"manufacturerLink,omitempty"`
	Name              string         `json:"name,omitempty"`
	Path              string         `json:"path,omitempty"`
	PaymentMethods    *PaymentMethod `json:"paymentMethods,omitempty"`
	SupportLink       string         `json:"supportLink,omitempty"`
	Translated        *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
	UpgradeVersion    string         `json:"upgradeVersion,omitempty"`
	UpgradedAt        *time.Time     `json:"upgradedAt,omitempty"`
	Version           string         `json:"version,omitempty"`
}
