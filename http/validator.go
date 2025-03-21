package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator interface {
	Validate(c *fiber.Ctx) error
}

func Validate(c *fiber.Ctx, v interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := c.BodyParser(v); err != nil {
		return err
	}
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}

/*
Example : 
type UserValidator struct {
	Name string `validate:"required"`
}

func (v *UserValidator) Validate(c *fiber.Ctx) error {
	//if needed, you can add more validations here
	return Validate(c, v)
}
*/
