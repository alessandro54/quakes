package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/alessandro54/quakes/internal/database"
	"github.com/alessandro54/quakes/internal/globals"
	"github.com/alessandro54/quakes/internal/models"
	"github.com/alessandro54/quakes/internal/providers"
	"google.golang.org/api/iterator"
	"log"
	"sort"
	"sync"
	"time"
)

func ListEarthquakes() []models.Earthquake {
	ctx := context.Background()

	client := database.Client()

	iter := client.Collection("earthquakes").OrderBy("report_number", firestore.Desc).Documents(ctx)

	var earthquakes []models.Earthquake

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var earthquake models.Earthquake
		if err := doc.DataTo(&earthquake); err != nil {
			log.Fatalf("Failed to convert data: %v", err)
		}

		earthquakes = append(earthquakes, earthquake)
	}
	defer client.Close()

	return earthquakes
}

func GetLatestEarthquake() *models.Earthquake {
	ctx := context.Background()

	client := database.Client()

	// Query for the latest document based on the reportNumber field
	iter := client.Collection("earthquakes").OrderBy("report_number", firestore.Desc).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if errors.Is(err, iterator.Done) {
		return nil
	}
	if err != nil {
		return nil
	}

	var earthquake models.Earthquake

	if err := doc.DataTo(&earthquake); err != nil {
		return nil
	}

	defer client.Close()

	return &earthquake
}

func CheckNewEarthquake() error {
	var wg sync.WaitGroup
	var providerData []models.Earthquake
	var lastDbReportNumber = 0

	wg.Add(2)

	go func() {
		defer wg.Done()

		lastEarthquake := GetLatestEarthquake()

		if lastEarthquake == nil {
			return
		}

		lastDbReportNumber = lastEarthquake.ReportNumber
	}()

	go func() {
		defer wg.Done()
		loc, _ := time.LoadLocation("America/Lima")
		providerData = providers.ByYear(time.Now().In(loc).Year())
	}()

	wg.Wait()

	sort.Sort(models.ByReportNumber(providerData))

	if providerData[0].ReportNumber != lastDbReportNumber {
		fmt.Printf("Data is not synchronized\n")
		globals.Busy = true
		syncEarthquakes(providerData, lastDbReportNumber)
	} else {
		fmt.Printf("Data is synchronized\n")
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
		fmt.Printf("Count: %v\n", count)
	}

	globals.Busy = false
}

func createEarthquake(earthquake models.Earthquake) error {
	client := database.Client()
	ctx := context.Background()

	_, _, err := client.Collection("earthquakes").Add(ctx, earthquake)
	if err != nil {
		fmt.Printf("Error adding earthquake: %v\n", err)
		return err
	}

	fmt.Printf("Earthquake added successfully\n")

	defer client.Close()
	return nil
}
