package main

import (
	"github.com/gin-gonic/gin"
	"ozone.com/book-service/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	models.ConnectDatabase()
	DefineRoutes(r)
	r.Run()
}
