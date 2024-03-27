package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-app/internal/api/rest"
	"go-blog-app/internal/api/rest/responses"
	"go-blog-app/internal/dto"
	"go-blog-app/internal/repository"
	"go-blog-app/internal/service"
	"go-blog-app/pkg/utils/validator"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func SetupArticleRoutes(rh *rest.RestHandler) {
	app := rh.App

	repo := repository.NewArticleRepository(rh.DB)
	svc := service.NewArticleService(repo, rh.Auth)

	handler := &ArticleHandler{
		svc: svc,
	}

	//Public Endpoints
	pubRoutes := app.Group("/articles")
	pubRoutes.Get("/", handler.GetArticles)
	pubRoutes.Get("/:id", handler.GetArticle)

	//Private Routes
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Post("/create", handler.CreatePost)
	pvtRoutes.Put("/update/:id", handler.UpdateArticle)
	pvtRoutes.Delete("/delete/:id", handler.DeleteArticle)
}

// CreatePost creates a new article.
// @Summary Create a new article
// @Description Creates a new article with provided details
// @Tags Articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.CreateArticleDto true "Article creation details"
// @Success 200 {object} domain.Article "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /articles/create [post]
func (h *ArticleHandler) CreatePost(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	article := dto.CreateArticleDto{}
	err := ctx.BodyParser(&article)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	errValidate := validator.ValidateStruct(ctx.Context(), &article)
	if errValidate != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, errValidate.Error())
	}

	err = h.svc.CreateArticle(article, user.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "error on create article")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, "article created")
}

// GetArticles retrieves all articles.
// @Summary Retrieve all articles
// @Description Retrieves all articles available
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {array} domain.Article "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /articles/ [get]
func (h *ArticleHandler) GetArticles(ctx *fiber.Ctx) error {
	articles, err := h.svc.GetArticles()
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, articles)
}

// GetArticle retrieves a specific article by ID.
// @Summary Retrieve article by ID
// @Description Retrieves a specific article by its ID
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} domain.Article "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /articles/{id} [get]
func (h *ArticleHandler) GetArticle(ctx *fiber.Ctx) error {
	articleId := ctx.Params("id")
	id, err := strconv.Atoi(articleId)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "invalid article ID")
	}
	article, err := h.svc.GetArticle(uint(id))
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, article)
}

// UpdateArticle updates an existing article.
// @Summary Update an existing article
// @Description Updates an existing article with provided details
// @Tags Articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Article ID"
// @Param input body dto.UpdateArticleDto true "Article update details"
// @Success 200 {object} "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /articles/update/{id} [put]// UpdateArticle updates an existing article.
func (h *ArticleHandler) UpdateArticle(ctx *fiber.Ctx) error {
	articleId := ctx.Params("id")
	id, err := strconv.Atoi(articleId)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "invalid article ID")
	}

	article := dto.UpdateArticleDto{}
	err = ctx.BodyParser(&article)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "invalid article inputs")
	}

	errValidate := validator.ValidateStruct(ctx.Context(), &article)
	if errValidate != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, errValidate.Error())
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	err = h.svc.UpdateArticle(article, uint(id), user.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, "article updated")
}

// DeleteArticle deletes an existing article.
// @Summary Delete an existing article
// @Description Deletes an existing article by its ID
// @Tags Articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Article ID"
// @Success 200 {string} string "article has been deleted"
// @Failure 400 {string} string "Bad Request"
// @Router /articles/delete/{id} [delete]
func (h *ArticleHandler) DeleteArticle(ctx *fiber.Ctx) error {
	articleId := ctx.Params("id")
	id, err := strconv.Atoi(articleId)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "invalid article ID")
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	errResp := h.svc.DeleteArticle(uint(id), user.ID)
	if errResp != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, errResp.Error())
	}
	return responses.NewSuccessResponse(ctx, http.StatusOK, "article has been deleted")
}
