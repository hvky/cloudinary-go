package main

import (
	"cloudinary-go/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/file", controllers.FileUpload())
	router.POST("/remote", controllers.RemoteUpload())

	router.Run("localhost:6000")
}
