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

func GetFormLogin(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/login.html")
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

	flash := map[string]interface{} {
		"FlashStatus": sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession": models.UserData,
	}


	return tmpl.Execute(c.Response(), flash)
}

func GetLogin(c echo.Context) error {
	err := c.Request(). ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := models.Users{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email =$1", email).Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword)
	if err != nil {
		return utilities.RedirectWithMessage(c, "Email Incorrect!", false,  "/form-login")
	}

	// membandingkan password yg di db dengan yg di html
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return utilities.RedirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM MASA BERLAKU LOGIN
	sess.Values["message"] = "Login success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Username
	sess.Values["email"] = user.Email
	sess.Values["Id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(),  c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}