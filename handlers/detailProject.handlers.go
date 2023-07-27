package handlers

import (
	"context"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"perosnal-web/utilities"
	"strconv"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetDetailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	templ, err := template.ParseFiles("views/detailProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	ProjectDetail := models.AddedProject{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.Technologies, &ProjectDetail.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = utilities.CountDuration(ProjectDetail.StartDate, ProjectDetail.EndDate)

	if utilities.CheckValue(ProjectDetail.Technologies, "nodeCheck") {
		ProjectDetail.NodeJs = true
	}
	if utilities.CheckValue(ProjectDetail.Technologies, "reactCheck") {
		ProjectDetail.ReactJs = true
	}
	if utilities.CheckValue(ProjectDetail.Technologies, "goCheck") {
		ProjectDetail.Golang = true
	}
	if utilities.CheckValue(ProjectDetail.Technologies, "jsCheck") {
		ProjectDetail.JavaScript = true
	}

	data := map[string]interface{}{
		"id":      id,
		"Project": ProjectDetail,
		"startDate": ProjectDetail.StartDate.Format("2006-01-02"),
		"endDate": ProjectDetail.EndDate.Format("2006-01-02"),
		"DataSession": models.UserData,
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		models.UserData.IsLogin = false
	} else {
		models.UserData.IsLogin = sess.Values["isLogin"].(bool)
		models.UserData.Name = sess.Values["name"].(string)
	}


	return templ.Execute(c.Response(), data)
}