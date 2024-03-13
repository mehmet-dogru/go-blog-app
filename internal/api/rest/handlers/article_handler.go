package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-app/internal/api/rest"
	"go-blog-app/internal/api/rest/responses"
	"go-blog-app/internal/dto"
	"go-blog-app/internal/repository"
	"go-blog-app/internal/service"
	"net/http"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func SetupArticleRoutes(rh *rest.RestHandler) {
	app := rh.App

	repo := repository.NewArticleRepository(rh.DB)
	svc := service.NewArticleService(repo, rh.Auth)

	handler := ArticleHandler{
		svc: svc,
	}

	//Public Endpoints
	pubRoutes := app.Group("/articles")
	pubRoutes.Get("/", handler.GetArticles)

	//Private Routes
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Post("/create", handler.CreatePost)
}

func (h *ArticleHandler) CreatePost(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	article := dto.CreateArticleDto{}
	err := ctx.BodyParser(&article)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	arc, err := h.svc.CreateArticle(article, user.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "error on create article")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, arc)
}

func (h *ArticleHandler) GetArticles(ctx *fiber.Ctx) error {
	articles, err := h.svc.GetArticles()
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, articles)
}
