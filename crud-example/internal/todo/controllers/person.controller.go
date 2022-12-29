package controllers

import (
	"github.com/HoSySon1802/crud--golang.git/internal/todo/models"
	"github.com/HoSySon1802/crud--golang.git/internal/todo/service"
	"github.com/HoSySon1802/crud--golang.git/internal/todo/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPeople(c *gin.Context) {
	people, err := service.GetPeople()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	}

	c.JSON(http.StatusOK, gin.H{"data": people})
}

func CreatePerson(c *gin.Context) {
	var input models.CreatePerson
	err := validate.ValidateCreatePerson(&input, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.CreatePerson(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetPerson(c *gin.Context) {
	var input models.GetPerson
	err := validate.ValidateGetPerson()
}
