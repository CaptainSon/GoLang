package repository

import (
	"testGin/entity"
)

func GetPeople() ([]entity.Person, error) {
	var people []entity.Person
	err := entity.DB.Find(&people).Error
	if err != nil {
		return nil, err
	}
	return people, nil
}

func CreatePerson(person *entity.Person) error {
	err := entity.DB.Create(person).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPerson(id uint) (person *entity.Person, err error) {
	err = entity.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}

func UpdatePerson(id uint, person2 *entity.Person) (person *entity.Person, err error) {
	err = entity.DB.Model(&person).Where("id = ?", id).Updates(&person2).Error
	if err != nil {
		return nil, err
	}
	return person2, nil
}
