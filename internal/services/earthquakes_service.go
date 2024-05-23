package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/alessandro54/quakes/internal/database"
	"github.com/alessandro54/quakes/internal/models"
	"github.com/alessandro54/quakes/internal/providers"
	"google.golang.org/api/iterator"
	"strings"
	"sync"
	"time"
)

func ListEarthquakes() []models.Earthquake {
	var earthquakes []models.Earthquake

	ctx := context.Background()

	client := database.Client()

	iter := client.Collection("Earthquakes").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			break
		}

		var earthquake models.Earthquake
		if err := doc.DataTo(&earthquake); err != nil {

		}
		earthquakes = append(earthquakes, earthquake)
	}
	return earthquakes
}

func CheckNewEarthquake() error {
	var wg sync.WaitGroup
	var dbData, providerData []models.Earthquake

	wg.Add(2)

	go func() {
		defer wg.Done()
		dbData = ListEarthquakes()
	}()

	go func() {
		defer wg.Done()
		loc, _ := time.LoadLocation("America/Lima")
		providerData = providers.ByYear(time.Now().In(loc).Year())
	}()

	wg.Wait()

	if (len(providerData) - len(dbData)) != 0 {
		fmt.Printf("Data is not synchronized\n")
		println(len(syncEarthquakes(providerData, dbData)))
	}
	return nil
}

func syncEarthquakes(dbData, providerData []models.Earthquake) []models.Earthquake {
	fmt.Printf("Syncing data...\n")
	existingEarthquakes := make(map[string]bool)

	for _, eq := range dbData {
		key := generateEarthquakeKey(eq)
		existingEarthquakes[key] = true
	}

	var newEarthquakes []models.Earthquake
	for _, eq := range providerData {
		key := generateEarthquakeKey(eq)

		if !existingEarthquakes[key] {
			newEarthquakes = append(newEarthquakes, eq)
		}
	}

	return newEarthquakes
}

func generateEarthquakeKey(eq models.Earthquake) string {
	return fmt.Sprintf("%f-%f-%s-%s", eq.Latitude, eq.Longitude, strings.TrimSpace(eq.LocalDate), strings.TrimSpace(eq.LocalTime))
}

func CreateEarthquake(earthquake models.Earthquake) error {
	//_, _, err = client.Collection("Earthquakes").Add(ctx, models.Earthquake{
	//	LocalDate: "11:37:33",
	//	LocalTime: "11:37:33",
	//	Latitude:  -13.53,
	//	Longitude: -76.2,
	//	Magnitude: 3.6,
	//	Depth:     33,
	//	Intensity: "III Tambo De Mora",
	//	Reference: "8 km al S de Tambo De Mora, Chincha - Ica",
	//})
	fmt.Printf(earthquake.Reference)

	//_, _, err = database.Client().
	//	Collection("Earthquakes").
	//	Add(context.Background(), earthquake)

	fmt.Printf("Earthquake added successfully\n")
	return nil
}
