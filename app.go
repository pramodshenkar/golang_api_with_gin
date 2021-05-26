package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pramodshenkar/golang_todo_api/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		hello := new(controllers.HelloWorldController)
		v1.GET("/hello", hello.Default)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":5000")
}
