package main

import (
	"github.com/Kazukite12/Go-Gin-Api/controllers/productController"
	"github.com/Kazukite12/Go-Gin-Api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.GET("api/product", productController.Index)
	router.GET("api/product/:id", productController.Show)
	router.POST("api/product", productController.Create)
	router.PUT("api/product/:id", productController.Update)
	router.DELETE("api/product", productController.Delete)

	router.Run("localhost:8080")
}
