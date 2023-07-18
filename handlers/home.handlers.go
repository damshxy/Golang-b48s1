package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func GetAddedProject(c echo.Context) error {
	title := c.FormValue("projectName")
	description := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")

	fmt.Println("Title :" + title)
	fmt.Println("Description :" + description)
	fmt.Println("Start Date :" + startDate)
	fmt.Println("End Date :" + endDate)

	return c.Redirect(http.StatusMovedPermanently, "/")
}