package handlers

import (
	"context"
	"html/template"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"perosnal-web/utilities"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetAddedEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/editedProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail := models.AddedProject{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT tb_project.id, tb_user.username, tb_project.project_id, tb_name_project, tb_project.content, tb_project.technologies, tb_proejct.image, tb_project.start_date, tb_project.end_date FROM tb_project LEFT JOIN tb_user ON tb_project.project_id = tb_user.id", id).Scan(&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.Technologies, &ProjectDetail.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
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

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		models.UserData.IsLogin = false
	} else {
		models.UserData.IsLogin = sess.Values["isLogin"].(bool)
		models.UserData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{} {
		"id": id,
		"Project": ProjectDetail,
		"StartDate": ProjectDetail.StartDate.Format("2006-01-02"),
		"EndDate": ProjectDetail.EndDate.Format("2006-01-02"),
		"DataSession": models.UserData,
	}

	return tmpl.Execute(c.Response(), data)
}

func GetSubmitEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	projectName := c.FormValue("projectName")
	content := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	image := c.FormValue("input-image")
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")

	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name_project=$1, start_date=$2, end_date=$3, description=$4, technologies[1]=$5, technologies[2]=$6, technologies[3]=$7, technologies[4]=$8, image=$9 WHERE id=$10", projectName, start, end, content, nodeCheck, reactCheck, golangCheck, jsCheck, image, id)
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

	return c.Redirect(http.StatusMovedPermanently, "/")
}