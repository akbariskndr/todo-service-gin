package dto

type CreateTodoDto struct {
	Title string `json:"title" binding:"required"`
}
