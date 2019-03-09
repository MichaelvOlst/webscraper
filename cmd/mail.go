package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"

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

		if app.config.Email.From == "" {
			logrus.Error(`You need to specify the "From" emailadres in the .env file`)
			return
		}

		if app.config.Email.To == "" {
			logrus.Error(`You need to specify the "To" emailadres in the .env file`)
			return
		}

		if app.config.Email.Password == "" {
			logrus.Error(`You need to specify the "Password" field in the .env file`)
			return
		}

		email := email.New(app.config.Email)
		data := struct{}{}
		_, err := email.Send("Found a website", data)
		if err != nil {
			logrus.Errorf("Could not send e-mail: %v\n", err)
			return
		}
		fmt.Println("Mail sent")
	},
}
