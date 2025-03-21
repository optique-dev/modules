package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// please implement the Repository interface

type Http struct {
	handlers []Handler
	app      *fiber.App
	config   *config.Config
}

func NewHttp(app *fiber.App, config *config.Config) *Http {
	return &Http{
		handlers: []Handler{},
		app:      app,
	}
}

func (m *Http) Register(handler Handler) {
	m.handlers = append(m.handlers, handler)
}

func (m *Http) Bootstrap() error {
	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})
	app.Use(recover.New())
	for _, handler := range m.handlers {
		handler.Register(app)
	}

	log.Fatal(app.Listen(m.config.Http.Port))
}

func (m *Http) Stop() error {
	return nil
}
