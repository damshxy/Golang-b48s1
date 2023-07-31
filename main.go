package main

import (
	"perosnal-web/connection"
	"perosnal-web/handlers"
	"perosnal-web/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// connection to database
	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	// Get Routing
	e.GET("/", handlers.GetHome)
	e.GET("/contact", handlers.GetContact)
	e.GET("/myProject", handlers.GetMyProject)
	e.GET("/testimonial", handlers.GetTestimonial)
	e.GET("/detailProject/:id", handlers.GetDetailProject)
	e.GET("/editProject/:id", handlers.GetAddedEditProject)
	// Auth Login And Register Get
	e.GET("/form-register", handlers.GetFormRegister)
	e.GET("/form-login", handlers.GetFormLogin)

	// Post Routing
	e.POST("/addedProject", middleware.GetUploadFile(handlers.GetAddedProject))
	e.POST("/deleteProject/:id", handlers.GetDeleteProject)
	e.POST("/edit-project/:id", handlers.GetSubmitEditProject)
	e.POST("/logout", handlers.GetLogout)
	// Auth Login And Register Post
	e.POST("/register", handlers.GetRegister)
	e.POST("/login", handlers.GetLogin)

	e.Logger.Fatal(e.Start("localhost:5000"))
}