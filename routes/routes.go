package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/gst-billing/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/products", addProduct)
	authenticated.GET("/products/:prod", getProduct)
	authenticated.GET("/billing/:prod/:quantity", generateBill)
	authenticated.GET("/billing", getBillings)
}
