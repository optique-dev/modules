package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Http struct {
	listen_addr string
	app         *fiber.App
	handlers    []Handler
}

func NewHttp(config Config) (*Http, error) {
	return &Http{
		listen_addr: config.ListenAddr,
		app:         fiber.New(),
		handlers:    []Handler{},
	}, nil
}

func (m *Http) WithHandler(handler Handler) {
	m.handlers = append(m.handlers, handler)
}

func (m *Http) Ignite() error {
	m.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	m.app.Use(logger.New())

	for _, handler := range m.handlers {
		handler.Register(m.app)
	}

	return m.app.Listen(m.listen_addr)
}

func (m *Http) Stop() error {
	return m.app.Shutdown()
}
