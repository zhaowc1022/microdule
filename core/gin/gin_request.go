package microdule_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hihibug/microdule/v2/rest/validator"
)

type GinRequest struct {
	Context   *gin.Context
	Validator validator.Validator
}

func NewGinRequest(c *gin.Context, v validator.Validator) *GinRequest {
	return &GinRequest{Context: c, Validator: v}
}

func (g *GinRequest) GetVal(c any) error {
	return g.Context.ShouldBind(c)
}

func (g *GinRequest) GetVerifyVal(c any) error {
	err := g.Context.ShouldBind(c)
	if err != nil {
		return err
	}
	return g.Validator.FetchValidatorError(c)
}
