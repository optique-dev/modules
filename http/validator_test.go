package http

import (
	"bytes"
	httpclient "net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
}

type User struct {
	Name string `json:"name" validate:"required"`
}

type UserValidator struct {
	BaseValidator
	Name string `json:"name" validate:"required"`
}

func (u *UserValidator) Validate(c *fiber.Ctx) error {
	return u.BaseValidator.Validate(c, u)
}

func NewUserController() userController {
	return userController{}
}

func IgniteHTTP(t *testing.T, app Http) {
	go func() {
		if err := app.Ignite(); err != nil {
			t.Fatal(err)
		}
	}()
}

func (u userController) Register(app *fiber.App) {
	app.Post("/user", func(c *fiber.Ctx) error {
		e := new(UserValidator)
		if err := e.Validate(c); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.SendString("User registered")
	})
}

func InitValidatorHTTP() Http {
	user := NewUserController()
	app, err := NewHttp(Config{
		ListenAddr: ":8081",
	})
	if err != nil {
		app.Stop()
		panic(err)
	}
	app.WithHandler(user)
	return app
}

func TestValidator(t *testing.T) {
	app := InitValidatorHTTP()
	defer app.Stop()
	IgniteHTTP(t, app)
	resp, err := httpclient.Post("http://localhost:8081/user", "application/json", bytes.NewBuffer([]byte(`{"name":""}`)))
	if err != nil {
		app.Stop()
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		app.Stop()
		t.Fatal("Health check failed")
	}
}

func TestValidatorShouldSucceed(t *testing.T) {
	app := InitValidatorHTTP()
	defer app.Stop()
	IgniteHTTP(t, app)
	resp, err := httpclient.Post("http://localhost:8081/user", "application/json", bytes.NewBuffer([]byte(`{"name":"John"}`)))
	if err != nil {
		app.Stop()
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		app.Stop()
		t.Fatal("Health check failed")
	}
}
