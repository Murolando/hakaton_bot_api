package handler

import (
	"net/http"
	"strconv"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"
)

// @Summary Stop Database
// @Description stop postgresql on server
// @Tags   actions
// @Accept  json
// @Produce  json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/solutions/kill-base/ [post]
func (h *Handler) killBaseAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	sshObject := c.Keys["ssh"].(ent.SSH)

	result, err := srv.KillBase(&sshObject)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok", result)
}

// @Summary Lock table
// @Description stop writing in table
// @Tags   actions
// @Accept  json
// @Produce  json
// @Param   tableName path string true "table name"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/solutions/stop-table/ [post]
func (h *Handler) stopTableAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	tableName := c.Param("tableName")
	if tableName == "" {
		newErrorResponse(c, http.StatusInternalServerError, "invalid tableName")
		return
	}

	result, err := srv.StopBase(tableName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok", result)
}

// @Summary Run table
// @Description start writing in table
// @Tags   actions
// @Accept  json
// @Produce  json
// @Param   tableName path string true "table name"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/solutions/run-table/ [post]
func (h *Handler) runTableAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	tableName := c.Param("tableName")
	if tableName == "" {
		newErrorResponse(c, http.StatusInternalServerError, "invalid tableName")
		return
	}

	result, err := srv.StopBase(tableName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok", result)
}


// @Summary Restart database
// @Description restart database
// @Tags   actions
// @Accept  json
// @Produce  json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/solutions/restart/ [post]
func (h *Handler) restartBaseAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	sshObject := c.Keys["ssh"].(ent.SSH)
	err := srv.RestartBase(&sshObject)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "ok", true)
}

// @Summary stop database
// @Description stop database
// @Tags   actions
// @Accept  json
// @Produce  json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/solutions/restart/ [post]
func (h *Handler) killProcessAction(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	pID, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	result, err := srv.KillProcess(int64(pID))

	newResponse(c, "ok", result)
}
