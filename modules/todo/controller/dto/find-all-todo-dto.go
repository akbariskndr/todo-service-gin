package dto

type FindAllTodoDto struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
	Search string `form:"search"`
}
