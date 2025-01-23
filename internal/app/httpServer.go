package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-task/config"
)

type Server struct {
	Addr   string
	Router *http.ServeMux
}

func NewHttpServer(c *config.Config, router *http.ServeMux) *Server {
	addr := fmt.Sprintf("%s:%s", c.HttpConf.Host, c.HttpConf.HttpPort)
	log.WithFields(log.Fields{"Address": addr}).Info("HTTP Server created")

	return &Server{
		Addr:   addr,
		Router: router,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    s.Addr,
		Handler: s.Router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
