package responses

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	errResp := ErrorResponse{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
	}

	return ctx.Status(statusCode).JSON(errResp)
}
