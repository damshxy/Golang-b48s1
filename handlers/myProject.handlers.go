package handlers

import (
	"fmt"
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

func GetAddedProject(c echo.Context) error {
	title := c.FormValue("projectName")
	description := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")

	fmt.Println("Title :" + title)
	fmt.Println("Description :" + description)
	fmt.Println("Start Date :" + startDate)
	fmt.Println("End Date :" + endDate)
	fmt.Println("Node Check :" + nodeCheck)
	fmt.Println("React Check  :" + reactCheck)
	fmt.Println("Golang Check  :" + golangCheck)
	fmt.Println("Java Script Check  :" + jsCheck)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}