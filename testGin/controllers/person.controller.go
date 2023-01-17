package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testGin/dto"
	"testGin/service"
	"testGin/transformer"
)

func GetPeople(c *gin.Context) {
	people, err := service.GetPeople()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	}
	c.JSON(http.StatusOK, gin.H{"data": people})
}

func CreatePerson(c *gin.Context) {
	var input dto.CreatePerson

	err := transformer.ValidateCreatePerson(&input, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.CreatePerson(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetPerson(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	person, err := service.GetPerson(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": person})
}

func UpdatePerson(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input dto.UpdatePerson

	err := transformer.ValidateUpdatePerson(&input, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := service.UpdatePerson(uint(id), dto.Person{
		Name: input.Name,
		Age:  input.Age,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})

}
