package routes

import (
	"encoding/json"
	"fmt"
	"github.com/alessandro54/quakes/internal/services"
	"net/http"
)

func EarthquakeIndex(w http.ResponseWriter, r *http.Request) {
	services.CheckNewEarthquake()
	earthquakes := services.ListEarthquakes()

	marshalled, err := json.Marshal(earthquakes)

	if err != nil {
		http.Error(w, "Error marshalling earthquakes", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(marshalled))
}
