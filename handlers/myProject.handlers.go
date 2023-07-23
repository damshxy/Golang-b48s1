package handlers

import (
	"net/http"
	"perosnal-web/models"
	"text/template"
	"time"

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
	projectName := c.FormValue("projectName")
	content := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")

	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	newProject :=  models.AddedProject{
		ProjectName: projectName,
		Content: content,
		StartDate: start,
		EndDate: end,
		NodeJs: (nodeCheck == "nodeCheck"),
		ReactJs: (reactCheck == "reactCheck"),
		Golang: (golangCheck == "goCheck"),
		JavaScript: (jsCheck == "jsCheck"),
	}

	models.DataProject = append(models.DataProject, newProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}