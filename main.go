package main

import(
	"os"
	service "github.com/vitor-caetano/gogo-service/service"
)

func main() {
	port := os.Getenv("Port")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}