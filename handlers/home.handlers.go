package handlers

import (
	"context"
	"fmt"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"perosnal-web/utilities"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name_project, start_date, end_date, description, technologies, image FROM tb_project")

	models.DataProject = []models.AddedProject{}
	for data.Next() {
		var each = models.AddedProject{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
		}

		each.Duration = utilities.CountDuration(each.StartDate, each.EndDate)

		// Node Js
		if utilities.CheckValue(each.Technologies, "nodeCheck") {
			each.NodeJs = true
		}
		// React Js
		if utilities.CheckValue(each.Technologies, "reactCheck") {
			each.ReactJs = true
		}
		// Golang
		if utilities.CheckValue(each.Technologies, "goCheck") {
			each.Golang = true
		}
		// JavaScript
		if utilities.CheckValue(each.Technologies, "jsCheck") {
			each.JavaScript = true
		}

		models.DataProject = append(models.DataProject, each)

	}

	project := map[string]interface{} {
		"Project": models.DataProject,
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return tmpl.Execute(c.Response(), project)
}