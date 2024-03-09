package src

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "jwt-auth/docs"
	"jwt-auth/src/packages/auth"
)

type ServerHandler struct {
}

func (handler *ServerHandler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootGroup := router.Group("/api")
	{
		auth.CreateGroup(rootGroup)
	}

	return router
}
