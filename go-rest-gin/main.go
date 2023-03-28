package main

import (
	"github.com/gagassurya19/go-rest-gin/controllers/productController"
	"github.com/gagassurya19/go-rest-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	r.Run()
}
