package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pjw1702/go-restapi-gin/controllers/productcontroller"
	"github.com/pjw1702/go-restapi-gin/models"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/pjw/products", productcontroller.Index)
	r.GET("/api/pjw/product/:id", productcontroller.Show)
	r.POST("/api/pjw/product", productcontroller.Create)
	r.PUT("/api/pjw/product/:id", productcontroller.Update)
	r.DELETE("/api/pjw/product", productcontroller.Delete)

	r.Run()
}
