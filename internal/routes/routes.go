package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/earthquakes", EarthquakeIndex)
	mux.HandleFunc("/earthquakes/latest", LatestEarthquake)

	return mux
}
