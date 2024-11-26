package rest

import (
	"github.com/hihibug/microdule/v2/rest/request"
	"github.com/hihibug/microdule/v2/rest/response"
)

type Rest interface {
	Client() any
	Run() error
	Request(c any) request.Request
	Response(c any) response.Response
}
