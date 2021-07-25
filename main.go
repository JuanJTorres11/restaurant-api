package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/JuanJTorres11/restaurant-api/Controller"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Route("/load", func(r chi.Router) {
		r.Get("/", Controller.LoadData)
		r.Get("/{date}", Controller.LoadData)
	})
	r.Route("/buyers", func(r chi.Router) {
		r.Get("/", Controller.ListBuyers)
		r.Get("/{buyerID}", Controller.GetBuyer)
	})

	http.ListenAndServe(":8000", r)
}
