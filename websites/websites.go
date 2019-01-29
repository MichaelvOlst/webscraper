package websites

// Website model for the database
type Website struct {
	ID         int64       `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	URL        string      `json:"url,omitempty"`
	Holder     string      `json:"holder,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

// Attribute ...
type Attribute struct {
	ID        int64  `json:"id,omitempty"`
	WebsiteID int64  `json:"website_id,omitempty"`
	Search    string `json:"search,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}
