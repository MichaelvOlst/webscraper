package models

// Attribute ...
type Attribute struct {
	ID        int64  `json:"id,omitempty"`
	WebsiteID int64  `json:"website_id,omitempty"`
	Search    string `json:"search,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}
