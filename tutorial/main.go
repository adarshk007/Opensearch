package main

import (
	"fmt"
	"os"

	"github.com/adarshk007/tutorial/client"
	"github.com/adarshk007/tutorial/config"
	"github.com/adarshk007/tutorial/router"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	// Load config singleton
	config.LoadConfig(env)

	// Initialize OpenSearch singleton
	client.InitOpenSearch()

	// Setup Gin router
	r := router.SetupRouter()
	fmt.Printf("Starting server on port %d\n", config.Cfg.Server.Port)
	r.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port))
}
