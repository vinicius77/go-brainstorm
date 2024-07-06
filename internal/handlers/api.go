package handlers

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/vinicius77/go-brainstorm/internal/middleware"
)

// Public Function
func Handler(router *chi.Mux) {
	router.Use(chimiddleware.StripSlashes)
	router.Route("/account", func(r chi.Router) {

		router.Use(middleware.Authorization)

		r.Get("/coins", GetCoinsBalance)

	})
}
