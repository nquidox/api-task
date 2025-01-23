package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"test-task/config"
	"test-task/internal/app"
	"test-task/internal/logging"
)

func main() {
	c := config.NewConfig()
	logging.LogSetup(c.AppConfig.LogLevel)

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	application := app.NewApp(c)
	go application.Start()

	sig := <-interrupt
	log.WithField("type", sig).Info("terminating, close app")
	os.Exit(0)
}
