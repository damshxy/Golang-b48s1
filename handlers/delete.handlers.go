package handlers

import (
	"net/http"
	"perosnal-web/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDeleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	models.DataProject = append(models.DataProject[:id], models.DataProject[id+1:]...)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}