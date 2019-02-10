package cronjob

import (
	"fmt"
	"time"

	"michaelvanolst.nl/scraper/datastore"
)

type job struct {
	interval int
	task     func()
	quit     chan struct{}
}

// Cronjob ..
type Cronjob struct {
	db   datastore.Datastore
	jobs []*job
}

type timeTicker struct {
}

// New returns a new type cronjob
func New(db datastore.Datastore) *Cronjob {
	return &Cronjob{
		db: db,
	}
}

// AddJob adds a func to the job
func (c *Cronjob) AddJob(interval int, t func()) {
	c.jobs = append(c.jobs, &job{
		interval: interval,
		task:     t,
		quit:     make(chan struct{}, 1),
	})
}

// Start runs the cronjob
func (c *Cronjob) Start() {

	for _, j := range c.jobs {

		// go j.task()
		go func(j *job) {
			ticker := time.NewTicker(time.Duration(j.interval) * time.Second)

			for {
				select {
				case <-ticker.C:
					j.task()
				case <-j.quit:
					ticker.Stop()
					fmt.Printf("Stopping job \n")
					return
				}
			}

		}(j)

	}
}

// Close closes all the jobs
func (c *Cronjob) Close() {
	for _, j := range c.jobs {
		close(j.quit)
	}
}

// func runCronjob() {

// 	ticker := time.NewTicker(5 * time.Second)

// 	sigs := make(chan os.Signal, 1)
// 	done := make(chan struct{}, 1)

// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		sig := <-sigs
// 		fmt.Println()
// 		fmt.Println(sig)
// 		done <- struct{}{}
// 	}()

// 	go func() {
// 		// var count int = 0
// 		for {
// 			select {
// 			case <-ticker.C:
// 				websites.Scrape(app.database)
// 			case <-done:
// 				fmt.Printf("Done counting.. \n")
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// 	fmt.Println("awaiting signal")
// 	<-done
// 	fmt.Println("exiting")
// }
