package handler

import (
	"net/http"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/service"
	"github.com/gin-gonic/gin"
)

// @Summary get base stat
// @Description stop database
// @Tags   info
// @Accept  json
// @Produce  json
// @Param input body ent.Config true "config"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/metrics/base-stat/ [post]
func (h *Handler) baseStat(c *gin.Context) {
	srv := c.Keys["service"].(*service.Service)
	bdName, ok := c.Get("dbName")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "placeQueryParams not found")
		return
	}
	bd := bdName.(string)

	metrics, err := srv.Metrics.BaseStat(bd)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	newResponse(c, "metrics", metrics)

}

// @Summary get user databases info
// @Description stop database
// @Tags   info
// @Accept  json
// @Produce  json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router /api/metrics/database-list/ [get]
func (h *Handler) dataBaseList(c *gin.Context) {
	database := make([]*ent.Database, 0, 3)
	database = append(database, &ent.Database{
		Host:     "localhost",
		Port:     "5433",
		User:     "postgres_test",
		DBName:   "postgres_test",
		Password: "123",
	})
	database = append(database, &ent.Database{
		Host:     "test_database",
		Port:     "5432",
		User:     "postgres_test",
		DBName:   "postgres_test",
		Password: "123",
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
