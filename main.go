package main

import (
	"SpecmaticProducer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/products", controllers.CreateProduct)
	engine.GET("/products", controllers.GetProductsByQuery)
	engine.Run(":8080")
}
