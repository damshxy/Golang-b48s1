package handlers

import (
	"context"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetMyProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/myProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		models.UserData.IsLogin = false
	} else {
		models.UserData.IsLogin = sess.Values["isLogin"].(bool)
		models.UserData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{} {
		"DataSession": models.UserData,
	}

	return tmpl.Execute(c.Response(), data)
}

func GetAddedProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	content := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	image := c.FormValue("input-image")
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (name_project, start_date, end_date, description, technologies[1], technologies[2], technologies[3], technologies[4], image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", projectName, startDate, endDate, content, nodeCheck, reactCheck, golangCheck, jsCheck, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}