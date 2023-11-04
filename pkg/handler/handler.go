package handler

import (
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"

	_ "github.com/Murolando/hakaton_bot_api/docs"
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
		metrics := api.Group("/metrics")
		{
			metrics.GET("/base-stat",h.baseStat)
		}
	}
	return router
}
