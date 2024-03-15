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

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	repo := repository.NewUserRepository(rh.DB)
	svc := service.NewUserService(repo, rh.Auth, *rh.Redis)

	handler := &UserHandler{
		svc: svc,
	}

	//Public Endpoints
	pubRoutes := app.Group("/users")
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	//Private Routes
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/profile", handler.GetProfile)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "error on signup")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	userInfo := h.svc.Auth.GetCurrentUser(ctx)

	user, err := h.svc.GetProfile(userInfo.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "user not found")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, user)
}
