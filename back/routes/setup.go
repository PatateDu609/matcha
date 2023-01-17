package routes

import (
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Setup() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(log.NewRouterLogger(log.Logger))
	router.Use(middleware.Recoverer)
	router.Use(database.AcquireMiddleware)

	//goland:noinspection HttpUrlsUsage
	corsOptions := cors.Options{
		Debug:            true,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodOptions,
			http.MethodConnect,
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
			http.MethodPatch,
		},
	}

	router.Use(cors.New(corsOptions).Handler)

	router.Route("/user", func(r chi.Router) {
		r.Get("/", getCurrentUser) // returns the current user (based on its session id)
		r.Get("/{uuid}", getUser)  // returns the pointed out user

		r.Put("/verify", verify)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Post("/sign-up", signUp)
		r.Post("/log-in", logIn)
	})
	return
}
