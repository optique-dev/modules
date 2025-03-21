package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) {
	code := fiber.statusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
