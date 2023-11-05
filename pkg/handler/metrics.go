package handler

import (
	"net/http"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) baseStat(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	bdName := c.Param("bdName")
	if bdName == ""{
		newErrorResponse(c, http.StatusInternalServerError, "invalid bdName")
		return
	}

	metrics, err := srv.Metrics.BaseStat(bdName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	newResponse(c, "metrics", metrics)

}

func (h *Handler) dataBaseList(c *gin.Context) {
	database := make([]*ent.Database, 0, 3)
	database = append(database, &ent.Database{Host: "", Port: "", User: "", DBName: "", Password: ""})
	database = append(database, &ent.Database{
		Host:     "mock",
		Port:     "mock",
		User:     "mock",
		DBName:   "mock",
		Password: "mock",
	})
	database = append(database, &ent.Database{
		Host:     "mock",
		Port:     "mock",
		User:     "mock",
		DBName:   "mock",
		Password: "mock",
	})
	newResponse(c, "database-list", database)

}
