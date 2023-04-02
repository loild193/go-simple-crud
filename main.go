package main

import (
	"go-simple-crud-1/controllers"
	"go-simple-crud-1/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	router := gin.Default()

	router.POST("/", controllers.CreatePost)
	router.GET("/", controllers.GetPosts)
	router.GET("/:id", controllers.GetPost)
	router.PUT("/:id", controllers.UpdatePost)
	router.DELETE("/:id", controllers.DeletePost)

	router.Run()
}
