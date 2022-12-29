package service

import (
	"github.com/HoSySon1802/crud--golang.git/internal/todo/entity"
	"github.com/HoSySon1802/crud--golang.git/internal/todo/models"
	"github.com/HoSySon1802/crud--golang.git/internal/todo/repository"
)

func GetPeople() ([]entity.Person, error) {
	people, err := repository.GetPeoPle()

	if err != nil {
		return nil, err
	}
	return people, err
}

func CreatePerson(input models.CreatePerson) error {

	err := repository.CreatePerson(&entity.Person{
		Name: input.Name,
		Age:  input.Age,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetPerson(id uint) (*entity.Person, error) {
	person, err := repository.GetPerson(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}