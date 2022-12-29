package repository

import "github.com/HoSySon1802/crud--golang.git/internal/todo/entity"

func GetPeoPle() ([]entity.Person, error) {
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
	err = entity.DB.Where("id = ?", id).First(&person).Errorirst
	if err != nil {
		return nil, err
	}
	return person, nil
}
