package routes

import (
	log2 "github.com/PatateDu609/matcha/utils/log"
	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func Setup() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(log2.NewRouterLogger(log2.Logger))
	router.Use(middleware.Recoverer)

	return
}
