package handler

import (
	"github.com/gin-gonic/gin"

	_ "github.com/Murolando/hakaton_bot_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
			metrics.GET("/base-stat", h.makeConnection, h.baseStat)
			metrics.GET("/database-list", h.dataBaseList)
		}
		solutions := api.Group("/solutions")
		{
			solutions.POST("/kill-process/:pid", h.makeConnection,h.killProcessAction)
			solutions.POST("/kill-base", h.makeConnection,h.killBaseAction)
			solutions.POST("/restart", h.makeConnection, h.restartBaseAction)
			solutions.POST("/stop-table/:tableName", h.makeConnection, h.stopTableAction)
			solutions.POST("/run-table/:tableName", h.makeConnection, h.runTableAction)
			// solutions.GET("/user-move",)
		}
	}
	return router
}
