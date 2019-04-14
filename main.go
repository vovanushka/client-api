package main

import (
	"os"

	"github.com/vovanushka/client-api/server"
	"github.com/vovanushka/client-api/service"
)

func main() {
	// Create port service
	portService := service.PortService{}
	// Connect to port server
	portService.Connect(getEnv("PORT_SERVICE_URL", ":7777"))
	defer portService.Close()
	// Create new server
	s := server.NewServer(&portService)
	// Start new http server
	s.Start(getEnv("HTTP_PORT", "8080"))
}

// helper function to make easier setting default env parameters
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
