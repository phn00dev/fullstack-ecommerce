package response

import "github.com/gofiber/fiber/v2"

type SucessResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func Success(ctx *fiber.Ctx, statusCode int, message string, data any) error {
	return ctx.Status(statusCode).JSON(SucessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
