package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/cakes", app.GetAllCakes)
	mux.Get("/cakes/{id}", app.GetCakeByID)
	mux.Post("/cakes", app.CreateCake)
	mux.Patch("/cakes/{id}", app.UpdateCakeByID)
	mux.Delete("/cakes/{id}", app.DeleteCake)

	return mux
}
