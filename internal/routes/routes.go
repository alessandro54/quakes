package routes

import (
	"fmt"
	"net/http"
	"github.com/alessandro54/quakes/internal/services"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, services.ByYear(2024))
}
