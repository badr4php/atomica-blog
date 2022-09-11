package main

import (
	"atomica-blog/db"
	"atomica-blog/models"
	"atomica-blog/routes"
)

func main() {
	dbInstance, _ := db.Open()
	dbInstance.AutoMigrate(&models.Post{})
	routes.Start()
}
