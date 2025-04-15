package handler

import "github.com/gofiber/fiber/v2"

func ResponseHandler[T any](c *fiber.Ctx, code int, result T) error {
	return c.Status(code).JSON(fiber.Map{
		"data": result,
		"path": c.OriginalURL(),
	})
}
