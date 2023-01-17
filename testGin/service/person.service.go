package service

import (
	"testGin/dto"
	"testGin/entity"
	"testGin/repository"
)

func GetPeople() ([]entity.Person, error) {
	people, err := repository.GetPeople()

	if err != nil {
		return nil, err
	}
	return people, nil
}

func CreatePerson(input dto.CreatePerson) error {

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

func UpdatePerson(id uint, person dto.Person) (*entity.Person, error) {
	updatedPerson, err := repository.UpdatePerson(id, &entity.Person{
		Name: person.Name,
		Age:  person.Age,
	})
	updatedPerson.ID = id
	if err != nil {
		return nil, err
	}

	return updatedPerson, nil
}
