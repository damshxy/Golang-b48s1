package handlers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetContact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}