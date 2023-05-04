package external_api

import (
	"github.com/MihaTeam1/golang_example/internal/handler/external_api/auth"
	"github.com/MihaTeam1/golang_example/internal/service"
	"github.com/gin-gonic/gin"
)

func initRouter(s *service.Service) *gin.Engine {
	router := gin.New()
	authHandler := auth.NewHandler(s)
	v1 := router.Group("/v1")
	{
		authGroup := v1.Group("/user")
		{
			authGroup.GET("/sign-up", authHandler.SignUp)
			authGroup.POST("/sign-in", authHandler.SignUp)
		}
	}
	return router
}
