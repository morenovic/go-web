package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello Victor",
		})
	})
	// Corremos nuestr
	router.Run()
}