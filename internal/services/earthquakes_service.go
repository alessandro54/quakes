package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/alessandro54/quakes/internal/database"
	"github.com/alessandro54/quakes/internal/globals"
	"github.com/alessandro54/quakes/internal/models"
	"google.golang.org/api/iterator"
	"log"
	"net/url"
	"strconv"
	"time"
)

func ListEarthquakes(page int, limit int, params url.Values) []models.Earthquake {
	ctx := context.Background()

	client := database.Client()

	dbEarthquakes := listFilters(client.Collection("earthquakes"), params)

	if params.Get("all") != "true" {
		dbEarthquakes = dbEarthquakes.
			OrderBy("report_number", firestore.Desc).
			Offset((page - 1) * limit).
			Limit(limit)
	}

	iter := dbEarthquakes.Documents(ctx)

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

func LatestEarthquake() *models.Earthquake {
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

func listFilters(collection *firestore.CollectionRef, params url.Values) firestore.Query {
	startDate := params.Get("start_date")
	endDate := params.Get("end_date")
	maxMagnitude := params.Get("max-mag")
	minMagnitude := params.Get("min-mag")
	all := params.Get("all")

	earthquakes := collection.Query

	if startDate != "" {
		earthquakes = earthquakes.Where("local_date", ">=", startDate)
	} else {
		if all != "true" {
			minDate := fmt.Sprintf("%d-01-01", time.Now().In(globals.Location).Year())
			earthquakes = earthquakes.Where("local_date", ">=", minDate)
		}
	}

	if endDate != "" {
		earthquakes = earthquakes.Where("local_date", "<=", endDate)
	}

	if maxMagnitude != "" {
		if maxMag, err := strconv.ParseFloat(maxMagnitude, 64); err == nil {
			earthquakes = earthquakes.Where("magnitude", "<=", maxMag)
		}
	}

	if minMagnitude != "" {
		if minMag, err := strconv.ParseFloat(minMagnitude, 64); err == nil {
			earthquakes = earthquakes.Where("magnitude", ">=", minMag)
		}
	}

	return earthquakes
}
