package handlers

import (
	"net/http"
	"perosnal-web/models"
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
	projectName := c.FormValue("projectName")
	content := c.FormValue("input-description")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	nodeCheck := c.FormValue("nodeCheck")
	reactCheck := c.FormValue("reactCheck")
	golangCheck := c.FormValue("goCheck")
	jsCheck := c.FormValue("jsCheck")

	// fmt.Println("Title :" + title)
	// fmt.Println("Description :" + description)
	// fmt.Println("Start Date :" + startDate)
	// fmt.Println("End Date :" + endDate)
	// fmt.Println("Node Check :" + nodeCheck)
	// fmt.Println("React Check  :" + reactCheck)
	// fmt.Println("Golang Check  :" + golangCheck)
	// fmt.Println("Java Script Check  :" + jsCheck)

	newProject :=  models.AddedProject{
		ProjectName: projectName,
		Content: content,
		StartDate: startDate,
		EndDate: endDate,
		NodeJs: nodeCheck,
		ReactJs: reactCheck,
		Golang: golangCheck,
		JavaScript: jsCheck,
	}

	models.DataProject = append(models.DataProject, newProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}