package handlers

import (
	"context"
	"log"
	"net/http"
	"perosnal-web/connection"
	"perosnal-web/models"
	"perosnal-web/utilities"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetFormRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/register.html")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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

func GetRegister(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	inputUsername := c.FormValue("input-username")
	inputEmail := c.FormValue("input-email")
	inputPassword := c.FormValue("input-password")
	
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(inputPassword), 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (username, email, password) VALUES ($1, $2, $3)", inputUsername, inputEmail, hashedPassword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		utilities.RedirectWithMessage(c, "Register failed please try again.", false, "/form-register")
	}

	return utilities.RedirectWithMessage(c, "Register Success!", true, "/form-login")
}