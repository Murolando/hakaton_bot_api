package handler

import (
	"github.com/Murolando/default_rest_arch/pkg/service"
	"github.com/gin-gonic/gin"

	_ "github.com/Murolando/default_rest_arch/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// router.Static("/storage","./storage")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.GET("/refresh/:refresh", h.newRefresh)
		}
	}
	return router
}
