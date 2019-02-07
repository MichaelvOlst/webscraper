package models

// Link is the data structure
type Link struct {
	ID           int64  `json:"id,omitempty"`
	WebsiteID    int64  `json:"website_id,omitempty"`
	URL          string `json:"url,omitempty"`
	Price        string `json:"price,omitempty"`
	Address      string `json:"address,omitempty"`
	Status       string `json:"status,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	ModifiedDate string `json:"modified_date,omitempty"`
	CreatedDate  string `json:"created_date,omitempty"`
}
