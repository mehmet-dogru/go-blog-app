package responses

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func NewSuccessResponse(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	successResp := SuccessResponse{
		Success:    true,
		StatusCode: statusCode,
		Data:       data,
	}

	return ctx.Status(statusCode).JSON(successResp)
}
