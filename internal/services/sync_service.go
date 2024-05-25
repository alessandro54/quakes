package services

import (
	"fmt"
	"github.com/alessandro54/quakes/internal/globals"
	"github.com/alessandro54/quakes/internal/models"
	"github.com/alessandro54/quakes/internal/providers"
	"log"
	"sort"
	"sync"
)

func DetectNewEarthquakes() error {
	var wg sync.WaitGroup
	var providerData []models.Earthquake
	var lastDbReportNumber = 0

	wg.Add(2)

	go func() {
		defer wg.Done()

		lastEarthquake := LatestEarthquake()

		if lastEarthquake == nil {
			return
		}

		lastDbReportNumber = lastEarthquake.ReportNumber
	}()

	go func() {
		defer wg.Done()
		//loc, _ := time.LoadLocation("America/Lima")
		providerData = providers.ByYear(2022)
	}()

	wg.Wait()

	sort.Sort(models.ByReportNumber(providerData))

	if providerData[0].ReportNumber != lastDbReportNumber {
		log.Println("SYNC SERVICE | DATA IS NOT SYNCHRONIZED")
		globals.Busy = true
		syncEarthquakes(providerData, lastDbReportNumber)
	} else {
		log.Println("SYNC SERVICE | DATA IS SYNCHRONIZED")
	}
	return nil
}

func syncEarthquakes(providerData []models.Earthquake, latestDBNumber int) {
	latestProviderNumber := providerData[0].ReportNumber

	toSync := latestProviderNumber - latestDBNumber
	count := toSync

	for i := 0; i < toSync; i++ {
		err := createEarthquake(providerData[i])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		count--
		log.Printf("SYNC SERVICE | Syncing data... %v\n", count)
	}

	globals.Busy = false
}
