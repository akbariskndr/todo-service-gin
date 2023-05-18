package dto

type UpdateTodoDto struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}
