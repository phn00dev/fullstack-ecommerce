package response

import "github.com/gofiber/fiber/v2"

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Errors     any    `json:"errors"`
}

func Error(ctx *fiber.Ctx, statusCode int, message string, errors any) error {
	return ctx.Status(statusCode).JSON(errorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	})
}
