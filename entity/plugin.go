package entity

import "time"

// completed

type Plugin struct {
	Active bool   `json:"active,omitempty"`
	Author string `json:"author,omitempty"`
	// Autoload map[type:object] `json:"autoload,omitempty"`
	BaseClass string `json:"baseClass,omitempty"`
	// Changelog map[type:object] `json:"changelog,omitempty"`
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
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	UpgradeVersion string     `json:"upgradeVersion,omitempty"`
	UpgradedAt     *time.Time `json:"upgradedAt,omitempty"`
	Version        string     `json:"version,omitempty"`
}
