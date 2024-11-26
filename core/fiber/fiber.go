package microdule_fiber

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/hihibug/microdule/v2/core/utils"
	httpConf "github.com/hihibug/microdule/v2/rest/config"
	"github.com/hihibug/microdule/v2/rest/request"
	"github.com/hihibug/microdule/v2/rest/response"
	"github.com/hihibug/microdule/v2/rest/validator"
)

type Fiber struct {
	Route     *fiber.App
	Config    *httpConf.Config
	Validator validator.Validator
}

func NewFiber(conf *httpConf.Config) *Fiber {
	defPath, _ := os.Getwd()
	fc := fiber.Config{
		DisableStartupMessage: true,
	}
	if conf.UseHtml {
		fc.Views = html.New(defPath+"/"+conf.TmpPath, ".html")
	}

	app := fiber.New(fc)

	// 请求日志
	path := defPath + "/" + conf.LogPath
	if ok, _ := utils.PathExists(path); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(path, os.ModePerm)
	}
	accessLogPath := path + "/access-" + time.Now().Format("2006-01-02") + ".log"
	// 记录到文件。
	f, _ := os.Create(accessLogPath)
	app.Use(logger.New(logger.Config{
		Output:     f,
		Format:     "[Fiber] ${time} | ${status} |  ${latency} | ${ip} | ${method}  ${path} \n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))

	// 初始化页面
	if conf.UseHtml {
		app.Static(defPath+"/"+conf.StaticPath, defPath+"/"+conf.TmpPath)
	}

	return &Fiber{
		Route:     app,
		Config:    conf,
		Validator: validator.NewValidator("zh"),
	}
}

func (f *Fiber) Client() any {
	return f
}

func (f *Fiber) Run() error {
	log.Printf("rest  port: %s \n", f.Config.Addr)
	return f.Route.Listen(":" + f.Config.Addr)
}

func (f *Fiber) Response(c any) response.Response {
	return NewFiberResponse(c.(*fiber.Ctx))
}

func (f *Fiber) Request(c any) request.Request {
	return NewFiberRequest(c.(*fiber.Ctx), f.Validator)
}
