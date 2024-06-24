package main

import (
	"rafaelnuansa/backend-api/controllers"
	"rafaelnuansa/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	router := gin.Default()
	// panggil function koneksi database di model
	models.ConnectDatabase()
	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {
		//return response JSON
		c.JSON(200, gin.H{
			"message": "Welcome to Go Programming Language, this project for learning Go-Lang with Mysql!",
		})
	})

	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/posts", controllers.StorePost)
	router.GET("/api/posts/:id", controllers.FindByPostId)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)
	//mulai server dengan port 3000
	router.Run(":3000")
}
