package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/gst-billing/db"
	"github.com/janeefajjustin/gst-billing/routes"
)

func main() {

	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
