package main

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}

//context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
