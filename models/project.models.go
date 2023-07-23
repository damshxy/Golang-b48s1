package models

import "time"

type AddedProject struct {
	Id          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	Duration string
	Content     string
	Technologies []string
	NodeJs      bool
	ReactJs     bool
	Golang      bool
	JavaScript bool
	Image string
}

var DataProject = []AddedProject{
	// {
	// 	ProjectName: "Bootcamp Dumways",
	// 	StartDate: "20/07/2023",
	// 	EndDate: "20/09/2023",
	// 	Content: "Dumbways",
	// 	NodeJs: "true",
	// 	ReactJs: "false",
	// 	Golang: "true",
	// 	JavaScript: "true",
	// },
}