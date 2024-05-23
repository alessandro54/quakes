package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/earthquakes", EarthquakeIndex)

	return mux
}
