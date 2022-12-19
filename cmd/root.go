package cmd

import (
	"fmt"
	"log"

	"gotoko-pos-api/internal/pkg/env"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "delivery service",
	Short: "Welcome to the delivery service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the delivery service.")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	env.Load()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
