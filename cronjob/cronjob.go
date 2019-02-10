package cronjob

import (
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
