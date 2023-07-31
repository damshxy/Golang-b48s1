package handlers

import (
	"context"
	"fmt"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"perosnal-web/utilities"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_user.username, tb_project.project_id, tb_name_project, tb_project.description, tb_project.technologies, tb_proejct.image, tb_project.start_date, tb_project.end_date FROM tb_project LEFT JOIN tb_user ON tb_project.project_id = tb_user.id")

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

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		models.UserData.IsLogin = false
	} else {
		models.UserData.IsLogin = sess.Values["isLogin"].(bool)
		models.UserData.Name = sess.Values["name"].(string)
	}

	project := map[string]interface{} {
		"Project": models.DataProject,
		"FlashStatus": sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession": models.UserData,
	}

	delete(sess.Values, "status")
	delete(sess.Values, "message")

	return tmpl.Execute(c.Response(), project)
}