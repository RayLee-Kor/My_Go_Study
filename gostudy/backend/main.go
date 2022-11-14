package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()     //loading env
		infrastructure.NewDatabase() //new database connection
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	router.Run(":8000")
}
