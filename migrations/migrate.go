package main

import (
	"go-simple-crud-1/initializers"
	"go-simple-crud-1/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
