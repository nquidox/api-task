package app

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-task/config"
	"test-task/internal/user"
	"test-task/pkg/repository"
)

type App struct {
	HttpServer *Server
}

func NewApp(c *config.Config) *App {
	router := http.NewServeMux()
	defaultRoute(router)

	database, err := repository.Connection(c)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Database connection failed")
	}

	user.NewUserHandler(database, router)

	return &App{
		HttpServer: NewHttpServer(c, router),
	}
}

func (a *App) Start() {
	if err := a.HttpServer.Run(); err != nil {
		log.WithField("error", err).Fatal("HTTP server")
	}
	log.Info("HTTP server started")
}
