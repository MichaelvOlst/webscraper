package cmd

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/spf13/cobra"
	"michaelvanolst.nl/scraper/email"
)

func init() {
	rootCmd.AddCommand(mailCmd)
}

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Print the version number of Scraper",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(app.config.Email)
		// fmt.Println(app.config.Email.From)
		// fmt.Println(app.config.Email.To)
		// fmt.Println(app.config.Email.Password)
		send("test from go", app.config.Email)
	},
}

func send(body string, cfg *email.Config) {

	from := cfg.From
	pass := cfg.Password
	to := cfg.To

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("email sent")
}
