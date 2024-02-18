package server

import (
	"go_notifier/configs"
	"go_notifier/internal/campaign"
	"go_notifier/internal/device"
	"go_notifier/internal/user"
	"go_notifier/internal/user_campaign"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	router              *chi.Mux
	port                string
	logger              log.FieldLogger
	deviceHandler       *device.DeviceHandler
	campaignHandler     *campaign.CampaignHandler
	userHandler         *user.UserHandler
	userCampaignHandler *user_campaign.UserCampaignHandler
}

func InitServer(
	serverConfig *configs.HttpServerConfig,
	logger log.FieldLogger,
	deviceHandler *device.DeviceHandler,
	campaignHandler *campaign.CampaignHandler,
	userHandler *user.UserHandler,
	userCampaignHandler *user_campaign.UserCampaignHandler,
) *HttpServer {
	return &HttpServer{
		router:              chi.NewRouter(),
		port:                serverConfig.ServerPort,
		logger:              logger,
		deviceHandler:       deviceHandler,
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

	s.router.Route("/device", func(r chi.Router) {
		r.Post("/", s.deviceHandler.CreateDevice)
	})

	s.router.Route("/campaign", func(r chi.Router) {
		r.Post("/", s.campaignHandler.CreateCampaign)
	})

	s.router.Route("/user-campaign", func(r chi.Router) {
		r.Post("/", s.userCampaignHandler.CreateUserCampaign)
	})
}
