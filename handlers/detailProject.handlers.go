package handlers

import (
	"net/http"
	"perosnal-web/models"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetDetailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ProjectDetail := models.AddedProject{}

	for i, data := range models.DataProject {
		if i == id {
			ProjectDetail = models.AddedProject{
				ProjectName: data.ProjectName,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Content: data.Content,
				NodeJs: data.NodeJs,
				ReactJs: data.ReactJs,
				Golang: data.Golang,
				JavaScript: data.JavaScript,
			}
		}
	}

	data := map[string]interface{}{
		"id":      id,
		"Project": ProjectDetail,
	}

	templ, err := template.ParseFiles("views/detailProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return templ.Execute(c.Response(), data)
}