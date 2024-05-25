package main

import (
	"fmt"
	"github.com/alessandro54/quakes/internal/globals"
	"github.com/alessandro54/quakes/internal/routes"
	"github.com/alessandro54/quakes/internal/services"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
)

func main() {
	c := cron.New()

	scheduleEarthquakeJob(c)

	c.Start()

	defer c.Stop()

	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on %s...\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func scheduleEarthquakeJob(c *cron.Cron) {
	job := func() {
		if !globals.Busy {
			services.DetectNewEarthquakes()
		}

	}
	_, err := c.AddFunc("@every 5m", job)

	if err != nil {
		log.Fatalf("Error scheduling job: %v\n", err)
	}
}
