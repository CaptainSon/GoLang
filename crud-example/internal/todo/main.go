package main

import (
	"github.com/HoSySon1802/crud--golang.git/internal/todo/controllers"
	"github.com/HoSySon1802/crud--golang.git/internal/todo/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	entity.ConnectDatabase() // new
	r.GET("/people", controllers.GetPeople)
	r.POST("/person", controllers.CreatePerson)
	//r.GET("/person/:id", controllers.GetPerson)
	//r.PUT("/person/:id", controllers.UpdatePerson)
	//
	r.Run("localhost:8080")
}
