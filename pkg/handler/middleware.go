package handler

import (
	"fmt"
	"net/http"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/repository"
	"github.com/Murolando/hakaton_bot_api/pkg/repository/postgres"
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) makeConnection(c *gin.Context) {
	var config ent.Config
	if err := c.BindJSON(&config); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(config)
	db, err := postgres.NewPostgresDB(&config)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ssh := ent.SSH{
		SSHUser:     config.SSHUser,
		SSHPassword: config.SSHPassword,
		Server:      config.Server,
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	c.Set("ssh", ssh)
	c.Set("dbName",config.DBName)
	c.Set("service", service)
}
