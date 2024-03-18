package dto

type CreateArticleDto struct {
	Title   string `json:"title" validate:"required,min=1,max=50"`
	Content string `json:"content" validate:"min=1,max=1000"`
}

type UpdateArticleDto struct {
	Title   string `json:"title" validate:"required,min=1,max=50"`
	Content string `json:"content" validate:"min=1,max=1000"`
}
