package main

import (
	"github.com/gin-gonic/gin"
	"testGin/controllers"
	"testGin/entity"
)

func main() {
	r := gin.Default()

	entity.ConnectDatabase() // new
	r.GET("/people", controllers.GetPeople)
	r.POST("/person", controllers.CreatePerson)
	r.GET("/person/:id", controllers.GetPerson)
	r.PUT("/person/:id", controllers.UpdatePerson)

	r.Run("localhost:8080")
}
