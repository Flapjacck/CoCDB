package main

import (
	"net/http"

	"github.com/flapjacck/CoCDB/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/api/data/*", routes.GetDataHandler)
	http.ListenAndServe(":3000", r)
}
