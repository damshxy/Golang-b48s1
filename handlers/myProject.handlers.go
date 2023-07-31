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
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")
	technologies := []string{nodeCheck, reactCheck, golangCheck, jsCheck}

	imageUpload := c.Get("dataFile").(string)
	
	sess, _ := session.Get("session", c)
	userID, _ := sess.Values["id"].(int)

	_, err := connection.Conn.Exec(context.Background(),
	"INSERT INTO tb_project (name_project, start_date, end_date, technologies, description, image, project_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", 
	projectName, startDate, endDate, technologies, content, imageUpload, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}
