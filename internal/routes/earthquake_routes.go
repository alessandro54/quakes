package routes

import (
	"encoding/json"
	"github.com/alessandro54/quakes/internal/models"
	"github.com/alessandro54/quakes/internal/services"
	"net/http"
	"net/url"
	"strconv"
)

type IndexResponse struct {
	Count int                 `json:"count"`
	Data  []models.Earthquake `json:"data"`
}

func EarthquakeIndex(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	page, limit := earthquakePagination(params)

	earthquakes := services.ListEarthquakes(page, limit, params)

	response := IndexResponse{
		Data:  earthquakes,
		Count: len(earthquakes),
	}

	marshalled, err := json.MarshalIndent(response, "", " ")

	if err != nil {
		http.Error(w, "Error marshalling earthquakes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func LatestEarthquake(w http.ResponseWriter, r *http.Request) {
	marshalled, err := json.MarshalIndent(services.LatestEarthquake(), "", " ")

	if err != nil {
		http.Error(w, "Error marshalling earthquake", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func earthquakePagination(params url.Values) (int, int) {
	pageParam := params.Get("page")
	limitParam := params.Get("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 5
	}

	return page, limit
}
