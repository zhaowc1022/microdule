package microdule_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hihibug/microdule/v2/rest/validator"
)

type FiberRequest struct {
	Context   *fiber.Ctx
	Validator validator.Validator
}

func (f FiberRequest) GetVal(c any) error {
	if f.Context.Method() == "GET" {
		return f.Context.QueryParser(c)
	}
	return f.Context.BodyParser(c)
}

func (f FiberRequest) GetVerifyVal(c any) error {
	var err error
	if f.Context.Method() == "GET" {
		err = f.Context.QueryParser(c)
	} else {
		err = f.Context.BodyParser(c)
	}
	if err != nil {
		return err
	}

	return f.Validator.FetchValidatorError(c)
}

func NewFiberRequest(c *fiber.Ctx, v validator.Validator) *FiberRequest {
	return &FiberRequest{Context: c, Validator: v}
}
