package main

import (
	"log"

	"clockwork/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "clockwork",
	Short: "Clockwork API server",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Clockwork API server",
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		v.AutomaticEnv() // Allow environment variable overrides

		// Read flags into viper
		port, _ := cmd.Flags().GetString("port")
		v.Set("server.port", port)

		log.Printf("Starting Clockwork API server on port %s...\n", port)
		server := api.NewClockworkApi(v)
		server.Serve()
	},
}

func init() {
	// Root command
	rootCmd.AddCommand(startCmd) // Add subcommands

	// Flag for start command
	startCmd.Flags().String("port", "50051", "Port to run the gRPC server on")
	viper.BindPFlag("server.port", startCmd.Flags().Lookup("port"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
