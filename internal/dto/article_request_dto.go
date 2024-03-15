package dto

type CreateArticleDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateArticleDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
