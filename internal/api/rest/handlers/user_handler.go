package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-app/internal/api/rest"
	"go-blog-app/internal/api/rest/responses"
	"net/http"
)

type UserHandler struct {
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := UserHandler{}

	pubRoutes := app.Group("/users")

	//Public Endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return responses.NewSuccessResponse(ctx, http.StatusOK, "Register")
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return responses.NewSuccessResponse(ctx, http.StatusOK, "Login")

}
