package repository

import (
	"errors"
	"go-blog-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type ArticleRepository interface {
	CreateArticle(article domain.Article) (domain.Article, error)
	FindArticleById(id uint) (domain.Article, error)
	UpdateArticle(id uint, u domain.Article) (domain.Article, error)
	GetArticles() ([]domain.Article, error)
	RemoveArticle(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r articleRepository) CreateArticle(article domain.Article) (domain.Article, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		log.Printf("create article error %v", err)
		return domain.Article{}, errors.New("failed to create article")
	}
	return article, nil
}

func (r articleRepository) FindArticleById(id uint) (domain.Article, error) {
	var article domain.Article

	err := r.db.First(&article, id).Error

	if err != nil {
		log.Printf("find article error %v", err)
		return domain.Article{}, errors.New("article does not exist")
	}

	return article, nil
}

func (r articleRepository) UpdateArticle(id uint, a domain.Article) (domain.Article, error) {
	var article domain.Article

	err := r.db.Model(&article).Clauses(clause.Returning{}).Where("id=?", id).Updates(a).Error

	if err != nil {
		log.Printf("error on update %v", err)
		return domain.Article{}, errors.New("failed update article")
	}

	return article, nil
}

func (r articleRepository) GetArticles() ([]domain.Article, error) {
	var articles []domain.Article
	err := r.db.Find(&articles).Error
	if err != nil {
		log.Printf("get articles error: %v", err)
		return nil, errors.New("failed to get articles")
	}
	return articles, nil
}

func (r articleRepository) RemoveArticle(id uint) error {
	var article domain.Article
	err := r.db.Where("id = ?", id).Delete(&article).Error

	if err != nil {
		log.Printf("error on delete %v", err)
		return errors.New("failed delete to article")
	}

	return err
}
