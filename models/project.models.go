package models

type AddedProject struct {
	ProjectName string
	StartDate string
	EndDate string
	Content string
	NodeJs string
	ReactJs string
	Golang string
	JavaScript string
}

var DataProject = []AddedProject {
	{
		ProjectName: "Bootcamp Dumways",
		StartDate: "20/07/2023",
		EndDate: "20/09/2023",
		Content: "sitLorem ipsum, dolor  amnsectetur aet codipisicing elit. Sint dolorum nostrum et non suscipit, veniam dignissimos unde error! Ducimus ipsum id officia suscipit quod libero omnis totam vitae eveniet iste. Lorem ipsum dolor sit amet consectetur, adipisicing elit. Aspernatur molestias cum voluptatibus ea necessitatibus dignissimos molestiae iste modi fugiat, nihil, consequatur in earum ex odit placeat dicta illo temporibus laudantium!",
		NodeJs: "true",
		ReactJs: "false",
		Golang: "true",
		JavaScript: "true",
	},
}