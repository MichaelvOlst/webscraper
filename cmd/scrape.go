package cmd

import (
	"github.com/spf13/cobra"
	"michaelvanolst.nl/scraper/websites"
)

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape the save urls",
	Run: func(cmd *cobra.Command, args []string) {

		websites.Scrape(app.database, app.config.Email)

	},
}
