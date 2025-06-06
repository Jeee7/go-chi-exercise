package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Rest API"))
	})

	return r
}
