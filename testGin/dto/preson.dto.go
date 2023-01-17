package dto

type Person struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type CreatePerson struct {
	Name string `json:"name" binding:"required"`
	Age  uint   `json:"age" binding:"required"`
}
type UpdatePerson struct {
	Name string `json:"name" binding:"required"`
	Age  uint   `json:"age" binding:"required"`
}

type GetPerson struct {
	ID uint `json:"id" binding:"required"`
}
