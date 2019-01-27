package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"michaelvanolst.nl/scraper/websites"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape the save urls",
	Run: func(cmd *cobra.Command, args []string) {
		urls, err := app.database.GetWebsites()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		// start := time.Now()
		ch := make(chan response, len(urls))
		defer close(ch)
		for _, url := range urls {
			go makeRequest(url, ch)
		}

		for r := range ch {

			fmt.Println(r.status + " - " + r.url)

		}

		// for range urls {
		// 	fmt.Println(<-ch)
		// }
		// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	},
}

type response struct {
	name   string
	url    string
	body   io.ReadCloser
	status string
	error  error
}

func makeRequest(url *websites.Website, ch chan<- response) {
	// start := time.Now()
	r, err := http.Get(url.URL)
	if err != nil {
		logrus.Error(err)
	}

	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".aanbodEntry").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		street := s.Find(".street-address").Text()
		link, _ := s.Find(".aanbodEntryLink").Attr("href")
		// title := s.Find("i").Text()
		fmt.Printf("house %d: %s --- %s\n", i, street, link)
	})

	// file, err := os.Create(url.Name + ".html")
	// if err != nil {
	// 	logrus.Error(err)
	// }
	// defer file.Close()

	// _, err = io.Copy(file, r.Body)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// secs := time.Since(start).Seconds()
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	logrus.Error(err)
	// }
	ch <- response{
		name:   url.Name,
		url:    url.URL,
		body:   r.Body,
		status: r.Status,
		error:  nil,
	}
}
