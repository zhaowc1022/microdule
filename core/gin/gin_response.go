package microdule_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hihibug/microdule/v2/core/utils"
	res "github.com/hihibug/microdule/v2/rest/response"
)

type (
	GinResponse struct {
		Context *gin.Context
	}
)

func NewGinResponse(c *gin.Context) *GinResponse {
	return &GinResponse{Context: c}
}

func GinResult(code int, data interface{}, msg string, c *gin.Context) error {
	data = utils.FmtLongInt(data)
	// 开始时间
	c.JSON(http.StatusOK, res.ResponseData{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	return nil
}

func (r *GinResponse) Ok() error {
	return GinResult(res.SUCCESS, map[string]interface{}{}, "success", r.Context)
}

func (r *GinResponse) OkWithMessage(message string) error {
	return GinResult(res.SUCCESS, map[string]interface{}{}, message, r.Context)
}

func (r *GinResponse) OkWithData(data interface{}) error {
	return GinResult(res.SUCCESS, data, "success", r.Context)
}

func (r *GinResponse) OkWithString(data string) error {
	r.Context.JSON(http.StatusOK, res.ResponseData{
		Code: res.SUCCESS,
		Data: data,
		Msg:  "success",
	})
	return nil
}

func (r *GinResponse) OkWithDetailed(data interface{}, message string) error {
	return GinResult(res.SUCCESS, data, message, r.Context)
}

func (r *GinResponse) Fail() error {
	return GinResult(res.ERROR, map[string]interface{}{}, "error", r.Context)
}

func (r *GinResponse) FailWithMessage(message string) error {
	return GinResult(res.ERROR, map[string]interface{}{}, message, r.Context)
}

func (r *GinResponse) FailWithDataMessages(err interface{}) error {
	return GinResult(res.ERROR, err, "参数校验错误", r.Context)
}

func (r *GinResponse) FailWithDetailed(data interface{}, message string) error {
	return GinResult(res.ERROR, data, message, r.Context)
}

func (r *GinResponse) JWTFailWithDetailed(data interface{}, message string) error {
	r.Context.JSON(http.StatusUnauthorized, res.ResponseData{
		Code: res.JWTERROR,
		Data: data,
		Msg:  message,
	})
	return nil
}

func (r *GinResponse) OkWithDataLogin(data interface{}, message string) error {
	r.Context.JSON(http.StatusOK, res.ResponseData{
		Code: res.JWTERROR,
		Data: data,
		Msg:  message,
	})
	return nil
}

func (r *GinResponse) OkStatusData(code int, data interface{}, message string) error {
	r.Context.JSON(http.StatusOK, res.ResponseData{
		Code: code,
		Data: data,
		Msg:  message,
	})
	return nil
}
