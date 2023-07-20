package main

import (
	"perosnal-web/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	// Get Routing
	e.GET("/", handlers.GetHome)
	e.GET("/contact", handlers.GetContact)
	e.GET("/myProject", handlers.GetMyProject)
	e.GET("/testimonial", handlers.GetTestimonial)
	e.GET("/detailProject/:id", handlers.GetDetailProject)
	// Post Routing
	e.POST("/addedProject", handlers.GetAddedProject)
	e.POST("/deleteProject/:id", handlers.GetDeleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}