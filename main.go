package main

import (
	"perosnal-web/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", handlers.GetHome)
	e.GET("/contact", handlers.GetContact)
	e.GET("/myProject", handlers.GetMyProject)
	e.GET("/testimonial", handlers.GetTestimonial)
	e.GET("/detailProject/:id", handlers.GetDetailProject)
	e.POST("/addedProject", handlers.GetAddedProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}