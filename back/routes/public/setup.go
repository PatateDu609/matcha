package public

import (
	"net/http"
	"net/http/httputil"

	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/routes/public/auth"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Setup() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(log.NewRouterLogger(log.Logger))
	router.Use(middleware.Recoverer)
	router.Use(database.AcquireMiddleware)

	//goland:noinspection HttpUrlsUsage
	corsOptions := cors.Options{
		Debug:            true,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:9000", "https://localhost:9000"},
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
	corsHandler := cors.New(corsOptions)
	corsHandler.Log = log.Logger
	router.Use(corsHandler.Handler)

	router.Handle("/socket.io/", httputil.NewSingleHostReverseProxy(config.Conf.SocketIO.ParsedURL))

	router.Route("/user", func(r chi.Router) {
		r.Get("/", getCurrentUser) // returns the current user (based on its session id)
		r.Get("/{uuid}", getUser)  // returns the pointed out user

		r.Get("/images/{uuid}", getImage)

		r.Put("/verify", verify)

		r.Patch("/set", setCurrentUser) // modify the current user (based on its session id)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json", "multipart/form-data"))

		r.Post("/sign-up", signUp)
		r.Post("/log-in", logIn)
		r.Post("/upload", uploadFile)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Route("/google", func(r chi.Router) {
			r.Get("/", auth.Google)
			r.Put("/redirect", auth.GoogleRedirect)
		})
	})

	return
}
