package http

import (
	log "github.com/sirupsen/logrus"
	"go-boilerplate/internal/bootstrap"
	"go-boilerplate/internal/server"
)

// Start function handler starting http listener
func Start() {
	bootstrap.SetJSONFormatter()

	serve := server.NewHTTPServer()
	defer serve.Done()

	if err := serve.Run(); err != nil {
		log.Fatalf("error running http server %v", err.Error())
	}

	return
}
