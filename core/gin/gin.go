package microdule_gin

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hihibug/microdule/v2/core/utils"
	"github.com/hihibug/microdule/v2/rest/request"
	"github.com/hihibug/microdule/v2/rest/response"
	"github.com/hihibug/microdule/v2/rest/validator"

	"github.com/gin-gonic/gin"
	httpConf "github.com/hihibug/microdule/v2/rest/config"
	"github.com/hihibug/microdule_gin/middleware"
)

type Gin struct {
	Route     *gin.Engine
	Config    *httpConf.Config
	Validator validator.Validator
}

func NewGin(conf *httpConf.Config) *Gin {
	gin.SetMode(gin.ReleaseMode)

	// 日志写入
	defPath, _ := os.Getwd()
	path := defPath + "/" + conf.LogPath
	if ok, _ := utils.PathExists(path); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(path, os.ModePerm)
	}
	accessLogPath := path + "/access-" + time.Now().Format("2006-01-02") + ".log"
	// 记录到文件。
	f, _ := os.Create(accessLogPath)
	gin.DefaultWriter = io.MultiWriter(f)

	var route = gin.Default()

	// 初始化页面
	if conf.UseHtml {
		defPath, _ := os.Getwd()
		route.Delims("{{", "}}")
		route.Static(defPath+"/"+conf.StaticPath, defPath+"/"+conf.TmpPath)
		route.LoadHTMLGlob(defPath + "/" + conf.TmpPath + "/*")
	}

	//注册GinCors
	route.Use(middleware.GinCors(), middleware.GinErrorHttp)
	route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found", "code": http.StatusNotFound})
	})

	// 验证器

	return &Gin{
		route,
		conf,
		validator.NewValidator("zh"),
	}
}

func (g *Gin) Client() any {
	return g
}

func (g *Gin) Run() error {
	log.Printf("rest  port: %s \n", g.Config.Addr)
	s := &http.Server{
		Addr:           ":" + g.Config.Addr,
		Handler:        g.Route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (g *Gin) Response(c any) response.Response {
	return NewGinResponse(c.(*gin.Context))
}

func (g *Gin) Request(c any) request.Request {
	return NewGinRequest(c.(*gin.Context), g.Validator)
}
