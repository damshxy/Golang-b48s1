package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetDetailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"title":   "Bootcamp Dumbways",
		"content": "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Sint dolorum nostrum et non suscipit, veniam dignissimos unde error! Ducimus ipsum id officia suscipit quod libero omnis totam vitae eveniet iste. Lorem ipsum dolor sit amet consectetur, adipisicing elit. Aspernatur molestias cum voluptatibus ea necessitatibus dignissimos molestiae iste modi fugiat, nihil, consequatur in earum ex odit placeat dicta illo temporibus laudantium!",
	}

	templ, err := template.ParseFiles("views/detailProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return templ.Execute(c.Response(), data)
}