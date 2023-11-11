package server

import (
	"go_notifier/configs"
	"go_notifier/internal/http/handlers/device"
	"go_notifier/internal/http/handlers/user"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type HttpServer struct {
	router *chi.Mux
	port   string
}

func InitServer(serverConfig *configs.HttpServerConfig) *HttpServer {
	return &HttpServer{
		router: chi.NewRouter(),
		port:   serverConfig.ServerPort,
	}
}

func (s *HttpServer) Start() error {
	s.initializeRoutes()
	log.Println("starting http server at:" + s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}

func (s *HttpServer) initializeRoutes() {
	s.router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	s.router.Route("/user", func(r chi.Router) {
		r.Post("/", user.CreateUserHandler)
	})

	s.router.Route("/device", func(r chi.Router) {
		r.Post("/", device.CreateDeviceHandler)
	})
}
