package entity

type Person struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
