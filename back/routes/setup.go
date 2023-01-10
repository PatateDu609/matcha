package routes

import (
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func Setup() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(log.NewRouterLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Group(func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		r.Post("/sign-up", signUp)
	})
	return
}
