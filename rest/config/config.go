package http

type Config struct {
	Mode       string `json:"mode" yaml:"mode"`             // rest 模型 默认gin 支持 gin fiber
	LogPath    string `json:"LogPath" yaml:"LogPath"`       // 访问日志地址
	UseHtml    bool   `json:"UseHtml" yaml:"UseHtml"`       // html 页面开关
	StaticPath string `json:"StaticPath" yaml:"StaticPath"` // 静态文件地址
	TmpPath    string `json:"TmpPath" yaml:"TmpPath"`       // html模板地址
	Addr       string `json:"addr" yaml:"addr"`             // 启动地址
}
