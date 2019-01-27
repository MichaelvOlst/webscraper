package websites

// Website model for the database
type Website struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
