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

// Register registers a new user.
// @Summary Register a new user
// @Description Registers a new user with provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.UserSignup true "User signup details"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /users/register [post]
func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	errValidate := validator.ValidateStruct(ctx.Context(), &user)
	if errValidate != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, errValidate.Error())
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "error on signup")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

// Login logs in a user.
// @Summary Log in user
// @Description Logs in a user with provided credentials
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.UserLogin true "User login details"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /users/login [post]
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "please provide valid inputs")
	}

	errValidate := validator.ValidateStruct(ctx.Context(), &loginInput)
	if errValidate != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, errValidate.Error())
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, token)
}

// GetProfile retrieves user profile information.
// @Summary Get user profile
// @Description Retrieves profile information of the logged-in user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ProfileInfo "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	userInfo := h.svc.Auth.GetCurrentUser(ctx)

	user, err := h.svc.GetProfile(userInfo.ID)
	if err != nil {
		return responses.NewErrorResponse(ctx, http.StatusBadRequest, "user not found")
	}

	return responses.NewSuccessResponse(ctx, http.StatusOK, user)
}
