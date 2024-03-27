package service

import (
	"errors"
	"go-blog-app/internal/domain"
	"go-blog-app/internal/dto"
	"go-blog-app/internal/helper"
	"go-blog-app/internal/repository"
)

type ArticleService struct {
	Repo repository.ArticleRepository
	Auth helper.Auth
}

func NewArticleService(repo repository.ArticleRepository, auth helper.Auth) ArticleService {
	return ArticleService{
		Repo: repo,
		Auth: auth,
	}
}

func (s ArticleService) CreateArticle(input dto.CreateArticleDto, authId uint) error {
	err := s.Repo.CreateArticle(domain.Article{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: authId,
	})

	return err
}

func (s ArticleService) GetArticles() ([]domain.Article, error) {
	articles, err := s.Repo.GetArticles()
	if len(articles) < 1 {
		return []domain.Article{}, errors.New("article not found")
	}
	return articles, err
}

func (s ArticleService) GetArticle(id uint) (*domain.Article, error) {
	article, err := s.Repo.FindArticleById(id)

	return &article, err
}

func (s ArticleService) UpdateArticle(input dto.UpdateArticleDto, id uint, authId uint) error {
	err := s.Repo.UpdateArticle(id, authId, domain.Article{Title: input.Title, Content: input.Content})
	return err
}

func (s ArticleService) DeleteArticle(id uint, authId uint) error {
	err := s.Repo.RemoveArticle(id, authId)
	return err
}
