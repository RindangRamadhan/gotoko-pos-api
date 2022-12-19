package cmd

import (
	"gotoko-pos-api/internal"
	"gotoko-pos-api/internal/handlers"
	"gotoko-pos-api/internal/pkg/env"

	"gotoko-pos-api/common/logger"
	"gotoko-pos-api/common/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	// Init Logger
	lgr := logger.New()
	logger.SetGlobalLogger(lgr)
	defer lgr.Close()

	srv, err := server.NewGinHttpRouter(env.Get().ServicePort)
	if err != nil {
		panic(err)
	}

	container := internal.NewContainer()

	router := handlers.NewRouter(srv.Router, container)
	router.RegisterRouter()

	srv.Start(env.Get().ServicePort)
}
