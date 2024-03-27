package repository

import (
	"errors"
	"go-blog-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type ArticleRepository interface {
	CreateArticle(article domain.Article) error
	FindArticleById(id uint) (domain.Article, error)
	UpdateArticle(id uint, authorId uint, u domain.Article) error
	GetArticles() ([]domain.Article, error)
	RemoveArticle(id uint, authorId uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r articleRepository) CreateArticle(article domain.Article) error {
	result := r.db.Create(&article)
	if result.Error != nil {
		log.Printf("create article error %v", result.Error)
		return errors.New("failed to create article")
	}
	return nil
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

func (r articleRepository) UpdateArticle(id uint, authorId uint, a domain.Article) error {
	var article domain.Article

	result := r.db.Model(&article).Clauses(clause.Returning{}).Where("id=?", id).Where("author_id=?", authorId).Updates(a)

	if result.Error != nil {
		log.Printf("error on update %v", result.Error)
		return errors.New("failed update article")
	}

	if result.RowsAffected == 0 {
		return errors.New("article not found or you are not authorized")
	}

	return nil
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

func (r articleRepository) RemoveArticle(id uint, authorId uint) error {
	var article domain.Article
	result := r.db.Where("id = ?", id).Where("author_id=?", authorId).Delete(&article)

	if result.Error != nil {
		log.Printf("error on delete %v", result.Error)
		return errors.New("failed delete to article")
	}

	if result.RowsAffected == 0 {
		return errors.New("article not found or you are not authorized")
	}

	return result.Error
}
