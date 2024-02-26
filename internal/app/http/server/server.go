package server

import (
	"go_notifier/configs"
	"go_notifier/internal/campaign"
	"go_notifier/internal/user"
	"go_notifier/internal/user_app"
	"go_notifier/internal/user_campaign"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	router              *chi.Mux
	port                string
	logger              log.FieldLogger
	userAppHandler      *user_app.UserAppHandler
	campaignHandler     *campaign.CampaignHandler
	userHandler         *user.UserHandler
	userCampaignHandler *user_campaign.UserCampaignHandler
}

func InitServer(
	serverConfig *configs.HttpServerConfig,
	logger log.FieldLogger,
	userAppHandler *user_app.UserAppHandler,
	campaignHandler *campaign.CampaignHandler,
	userHandler *user.UserHandler,
	userCampaignHandler *user_campaign.UserCampaignHandler,
) *HttpServer {
	return &HttpServer{
		router:              chi.NewRouter(),
		port:                serverConfig.ServerPort,
		logger:              logger,
		userAppHandler:      userAppHandler,
		campaignHandler:     campaignHandler,
		userHandler:         userHandler,
		userCampaignHandler: userCampaignHandler,
	}
}

func (s *HttpServer) Start() error {
	s.initializeRoutes()
	s.logger.Info("starting http server at:" + s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}

func (s *HttpServer) initializeRoutes() {
	s.router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	s.router.Route("/user", func(r chi.Router) {
		r.Post("/", s.userHandler.CreateUserHandler)
	})

	s.router.Route("/user-app", func(r chi.Router) {
		r.Post("/", s.userAppHandler.CreateUserApp)
	})

	s.router.Route("/campaign", func(r chi.Router) {
		r.Post("/", s.campaignHandler.CreateCampaign)
	})

	s.router.Route("/user-campaign", func(r chi.Router) {
		r.Post("/", s.userCampaignHandler.CreateUserCampaign)
	})
}
