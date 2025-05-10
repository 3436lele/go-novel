package interfaces

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/3436lele/go-novel/backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r        *gin.Engine
	server   *http.Server
	shutdown bool
	err      chan error
	logger   *slog.Logger
}

func NewServer(config *config.Config) *Server {
	log := slog.With("domain", "interfaces")

	gin.SetMode(config.Server.Mode)

	r := gin.New()
	r.Use(
		gin.Recovery(),
	)

	// CORS
	corsConfig := cors.Config{
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{
			"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	if config.Server.Mode == "debug" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowAllOrigins = true
	}

	r.Use(cors.New(corsConfig))

	s := &Server{
		r: r,
		server: &http.Server{
			Handler: r,
			Addr:    config.Server.Host + ":" + config.Server.Port,
		},
		shutdown: false,
		err:      make(chan error),
		logger:   log,
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Info("Starting server", "address", s.server.Addr)
	return s.server.ListenAndServe()
}
