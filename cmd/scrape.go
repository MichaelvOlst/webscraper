package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

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

		var wg sync.WaitGroup

		wg.Add(len(urls))

		for _, url := range urls {
			go makeRequest(url, ch, &wg)
		}

		wg.Wait()
	},
}

type response struct {
	name   string
	url    string
	body   io.ReadCloser
	status string
	error  error
}

func makeRequest(w *websites.Website, ch chan<- response, wg *sync.WaitGroup) {
	// start := time.Now()
	defer wg.Done()
	r, err := http.Get(w.URL)
	if err != nil {
		logrus.Error(err)
	}
	defer r.Body.Close()

	pu, err := url.Parse(w.URL)
	if err != nil {
		logrus.Error(err)
	}

	wURL := fmt.Sprintf("%s://%s", pu.Scheme, pu.Host)

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(w.Holder).Each(func(i int, s *goquery.Selection) {

		status := s.Find(".objectstatusbanner").Text()
		if strings.Contains(strings.ToLower(status), "nieuw") {

			street := s.Find(".street-address").Text()
			link, _ := s.Find(".aanbodEntryLink").Attr("href")

			fmt.Printf("%s %s --- %s%s\n", street, status, wURL, link)
		}
	})

	ch <- response{
		name:   w.Name,
		url:    w.URL,
		body:   r.Body,
		status: r.Status,
		error:  nil,
	}
}
