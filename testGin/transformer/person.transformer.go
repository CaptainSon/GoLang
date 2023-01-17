package transformer

import (
	"github.com/gin-gonic/gin"
	"testGin/dto"
)

func ValidateCreatePerson(input *dto.CreatePerson, c *gin.Context) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}
	return nil
}
func ValidateUpdatePerson(input *dto.UpdatePerson, c *gin.Context) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}
	return nil
}
