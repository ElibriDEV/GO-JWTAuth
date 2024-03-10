package main

import (
	"jwt-auth/initializators"
	"jwt-auth/src"
	"log"
)

// @title           GO JWT-Auth
// @version         1.0

// @host      localhost:8000
// @BasePath  /api

// @securityDefinitions.apikey Bearer-Access
// @in header
// @name Access
// @description Type: Bearer YOUR_ACCESS_TOKEN

// @securityDefinitions.apikey Bearer-Refresh
// @in header
// @name Refresh
// @description Type: Bearer YOUR_REFRESH_TOKEN

func main() {
	handler := new(src.ServerHandler)
	server := new(src.Server)

	initializators.LoadEnv()
	initializators.MongoInit(initializators.Config.MongoURL)

	log.Print("Swagger docs: http://localhost:", initializators.Config.ApplicationPort, "/docs/index.html")

	if err := server.Run(initializators.Config.ApplicationPort, handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
