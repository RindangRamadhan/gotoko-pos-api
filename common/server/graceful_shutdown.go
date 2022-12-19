package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GracefulShutdown struct {
	httpServer *http.Server
}

func NewGracefulShutdown(handler http.Handler, address string) *GracefulShutdown {
	return &GracefulShutdown{
		httpServer: &http.Server{
			Handler:      handler,
			Addr:         address,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (g *GracefulShutdown) GracefullyShutdown() {
	go func() {
		if err := g.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	log.Println("server is running...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err.Error())
		os.Exit(1)
	}

	log.Println("Server stopped.")
}
