package routes

import (
	"encoding/json"
	"github.com/alessandro54/quakes/internal/services"
	"net/http"
)

func EarthquakeIndex(w http.ResponseWriter, r *http.Request) {
	marshalled, err := json.MarshalIndent(services.ListEarthquakes(), "", " ")

	if err != nil {
		http.Error(w, "Error marshalling earthquakes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func LatestEarthquake(w http.ResponseWriter, r *http.Request) {
	marshalled, err := json.MarshalIndent(services.GetLatestEarthquake(), "", " ")

	if err != nil {
		http.Error(w, "Error marshalling earthquake", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}
