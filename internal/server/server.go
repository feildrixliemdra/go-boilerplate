package server

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go-boilerplate/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server interface
type Server interface {
	Run() error
	Done()
	//Config() *appctx.Config
}

// handle using generic http.Handler so it can be easily
// to switch between router library like gin, echo or chi
type httpServer struct {
	config  *config.Config
	handler http.Handler
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) Server {
	return &httpServer{
		handler: handler,
		config:  cfg,
	}
}

func (h *httpServer) Run() error {
	var err error

	srv := &http.Server{
		Addr:         ":" + h.config.App.Port,
		Handler:      h.handler,
		ReadTimeout:  time.Duration(h.config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(h.config.App.WriteTimeout) * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Info("err listen and serve %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //3 second timeout to make sure all process finished
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	// catching ctx.Done(). timeout of 3 seconds.
	select {
	case <-ctx.Done():
		log.Info("server exiting")
	}

	return nil
}

func (h *httpServer) Done() {
	log.Info("service http stopped")
}
