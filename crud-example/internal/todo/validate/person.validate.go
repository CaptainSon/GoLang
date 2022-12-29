package validate

import (
	"github.com/HoSySon1802/crud--golang.git/internal/todo/models"
	"github.com/gin-gonic/gin"
)

func ValidateCreatePerson(input *models.CreatePerson, c *gin.Context) error {
	err := c.ShouldBindJSON(input)
	if err != nil {
		return err
	}
	return nil
}

func ValidateGetPerson(input *models.GetPerson, c *gin.Context) error {
	err := c.ShouldBindJSON(input)
	if err != nil {
		return err
	}
	return nil
}
