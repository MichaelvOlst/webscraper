package websites

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"michaelvanolst.nl/scraper/email"

	"michaelvanolst.nl/scraper/datastore"
	"michaelvanolst.nl/scraper/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

// Scrape scrapes the saved websites
func Scrape(db datastore.Datastore, emailCfg *email.Config) {
	urls, err := db.GetWebsites()

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Println("Scraping urls...")

	ch := make(chan response, len(urls))
	defer close(ch)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		logrus.Println("Scraping url " + url.URL)
		go makeRequest(url, ch, &wg, db, emailCfg)
	}

	wg.Wait()
	logrus.Println("Done scraping...")
}

type response struct {
	name   string
	url    string
	body   io.ReadCloser
	status string
	error  error
}

func makeRequest(w *models.Website, ch chan<- response, wg *sync.WaitGroup, db datastore.Datastore, emailCfg *email.Config) {
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

	websiteURL := fmt.Sprintf("%s://%s", pu.Scheme, pu.Host)

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(w.Holder).Each(func(i int, s *goquery.Selection) {

		attributes := make(map[string]string)
		for _, a := range w.Attributes {
			attributes[a.Type] = a.Search
		}

		status := s.Find(attributes["status"]).Text()
		if strings.Contains(strings.ToLower(status), attributes["statustext"]) {

			var l models.Link

			address := s.Find(attributes["address"]).Text()
			link, _ := s.Find(attributes["link"]).Attr("href")

			imageURL, _ := s.Find(attributes["image"]).Attr("src")

			price := s.Find(attributes["price"]).Text()
			price = strings.TrimSpace(price)

			l.WebsiteID = w.ID
			l.Address = address
			l.URL = fmt.Sprintf("%s%s", websiteURL, link)
			l.Price = price
			l.Status = status
			l.ImageURL = imageURL

			err = db.SaveLink(&l)
			if err == nil {
				// fmt.Printf("%s %s %s --- %s%s\n", address, status, price, websiteURL, link)
				fmt.Printf("Found a new house %s - %s \n", address, price)

				td := struct {
					Address  string
					URL      string
					ImageURL string
					Price    string
				}{
					Address:  l.Address,
					URL:      l.URL,
					ImageURL: l.ImageURL,
					Price:    l.Price,
				}
				e := email.New(emailCfg)
				_, err := e.Send("Found a new house", td)
				if err != nil {
					fmt.Printf("Error send e-mail %v\n", err)
				}
			}
			// else {
			// 	// fmt.Printf("%s %s %s --- %s%s\n", address, status, price, websiteURL, link)
			// 	fmt.Printf("Found a new house %s - %s \n", address, price)
			// }

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
