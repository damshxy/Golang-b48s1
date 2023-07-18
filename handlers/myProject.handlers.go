package handlers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetMyProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/myProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}