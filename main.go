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

// @securityDefinitions.apikey Bearer-Refresh
// @in header
// @name Refresh

func main() {
	handler := new(src.ServerHandler)
	server := new(src.Server)

	initializators.LoadEnv()
	initializators.MongoInit(initializators.Env.MongoURL)

	log.Print("Swagger docs: http://localhost:", initializators.Env.ApplicationPort, "/docs/index.html")

	if err := server.Run(initializators.Env.ApplicationPort, handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
