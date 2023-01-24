package private

import (
	"net/http"

	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/routes/private/internal/db/room"
	"github.com/PatateDu609/matcha/routes/private/internal/session"
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
		AllowedOrigins:   []string{config.Conf.SocketIO.URL}, // these routes are only accessible by socket.io
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

	router.Route("/session", func(r chi.Router) {
		r.Get("/uuid", session.GetUUID)
	})

	router.Route("/db", func(r chi.Router) {
		r.Route("/room", func(roomRoute chi.Router) {
			roomRoute.Post("/create", room.Create)
			roomRoute.Get("/{sid}", room.Get)
			roomRoute.Delete("/{sid}", room.Delete)
		})
	})

	return
}
