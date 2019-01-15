package websites

import "fmt"

// Website model for the database
type Website struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Scrape the saved urls
func Scrape() {
	fmt.Println("Scraping...........")
}
