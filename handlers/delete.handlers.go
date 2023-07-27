package handlers

import (
	"context"
	"net/http"
	"perosnal-web/connection"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDeleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}