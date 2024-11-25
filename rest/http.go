package rest

import (
	conf "github.com/hihibug/microdule/rest/config"
	"github.com/hihibug/microdule/rest/request"
	"github.com/hihibug/microdule/rest/response"
	"github.com/hihibug/microdule/rest/rest_fiber"
	"github.com/hihibug/microdule/rest/rest_gin"
)

type Rest interface {
	Client() any
	Run() error
	Request(c any) request.Request
	Response(c any) response.Response
}

func NewRest(config *conf.Config) Rest {
	switch config.Mode {
	case "gin":
		return rest_gin.NewGin(config)
	case "fiber":
		return rest_fiber.NewFiber(config)
	default:
		return rest_gin.NewGin(config)
	}
}
