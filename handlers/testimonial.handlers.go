package handlers

import (
	"net/http"
	"perosnal-web/models"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetTestimonial(c echo.Context) error {
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

	tmpl, err := template.ParseFiles("views/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}