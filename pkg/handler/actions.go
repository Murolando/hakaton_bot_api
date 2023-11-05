package handler

import (
	"net/http"
	"strconv"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"
)

// stop database
func (h *Handler) killBaseAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	sshObject := c.Keys["ssh"].(ent.SSH)

	result,err := srv.KillBase(&sshObject)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok",result)
}

// lock table
func (h *Handler) stopTableAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	tableName := c.Param("tableName")
	if tableName == ""{
		newErrorResponse(c, http.StatusInternalServerError, "invalid tableName")
		return
	}

	result,err := srv.StopBase(tableName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok",result)
}

// run table
func (h *Handler) runTableAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	tableName := c.Param("tableName")
	if tableName == ""{
		newErrorResponse(c, http.StatusInternalServerError, "invalid tableName")
		return
	}

	result,err := srv.StopBase(tableName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok",result)
}	

// checkpoint and restart base
func (h *Handler) restartBaseAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	sshObject := c.Keys["ssh"].(ent.SSH)
	err := srv.RestartBase(&sshObject)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok",true)
}

// stop database
func (h *Handler) killProcessAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	pID, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	result,err := srv.KillProcess(int64(pID))

	newResponse(c, "ok",result)
}
