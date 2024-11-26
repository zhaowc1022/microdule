package microdule_fiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/hihibug/microdule/v2/core/utils"
	res "github.com/hihibug/microdule/v2/rest/response"
)

type (
	GinResponse struct {
		Context *fiber.Ctx
	}
)

func NewFiberResponse(c *fiber.Ctx) *GinResponse {
	return &GinResponse{Context: c}
}

func FiberResult(code int, data interface{}, msg string, c *fiber.Ctx) error {
	data = utils.FmtLongInt(data)
	// 开始时间
	return c.Status(code).JSON(res.ResponseData{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func (r *GinResponse) Ok() error {
	return FiberResult(res.SUCCESS, map[string]interface{}{}, "success", r.Context)
}

func (r *GinResponse) OkWithMessage(message string) error {
	return FiberResult(res.SUCCESS, map[string]interface{}{}, message, r.Context)
}

func (r *GinResponse) OkWithData(data interface{}) error {
	return FiberResult(res.SUCCESS, data, "success", r.Context)
}

func (r *GinResponse) OkWithString(data string) error {
	return r.Context.Status(http.StatusOK).JSON(res.ResponseData{
		Code: res.SUCCESS,
		Data: data,
		Msg:  "success",
	})
}

func (r *GinResponse) OkWithDetailed(data interface{}, message string) error {
	return FiberResult(res.SUCCESS, data, message, r.Context)
}

func (r *GinResponse) Fail() error {
	return FiberResult(res.ERROR, map[string]interface{}{}, "error", r.Context)
}

func (r *GinResponse) FailWithMessage(message string) error {
	return FiberResult(res.ERROR, map[string]interface{}{}, message, r.Context)
}

func (r *GinResponse) FailWithDataMessages(err interface{}) error {
	return FiberResult(res.ERROR, err, "参数校验错误", r.Context)
}

func (r *GinResponse) FailWithDetailed(data interface{}, message string) error {
	return FiberResult(res.ERROR, data, message, r.Context)
}

func (r *GinResponse) JWTFailWithDetailed(data interface{}, message string) error {
	return r.Context.Status(http.StatusUnauthorized).JSON(res.ResponseData{
		Code: res.JWTERROR,
		Data: data,
		Msg:  message,
	})
}

func (r *GinResponse) OkWithDataLogin(data interface{}, message string) error {
	return r.Context.Status(http.StatusOK).JSON(res.ResponseData{
		Code: res.JWTERROR,
		Data: data,
		Msg:  message,
	})
}

func (r *GinResponse) OkStatusData(code int, data interface{}, message string) error {
	return r.Context.Status(http.StatusOK).JSON(res.ResponseData{
		Code: code,
		Data: data,
		Msg:  message,
	})
}
