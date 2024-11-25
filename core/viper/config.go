package viper

import (
	"github.com/hihibug/microdule/v2/core/etcd"
	"github.com/hihibug/microdule/v2/core/redis"
	"github.com/hihibug/microdule/v2/core/zap"
	http "github.com/hihibug/microdule/v2/rest/config"
	"github.com/hihibug/microdule/v2/rpc"
)

type Config struct {
	DB    DbConfig      `json:"db" yaml:"db"`
	Etcd  *etcd.Config  `json:"etcd" yaml:"etcd"`
	Redis *redis.Config `json:"redis" yaml:"redis"`
	Log   *zap.Config   `json:"log" yaml:"log"`
	Http  *http.Config  `json:"rest" yaml:"rest"`
	Rpc   *rpc.Config   `json:"rpc" yaml:"rpc"`
}

type DbConfig struct {
	DbType      string `json:"db-type" yaml:"dbType"`
	Path        string `json:"path" yaml:"path"`
	Config      string `json:"config" yaml:"config"`
	Dbname      string `json:"dbname" yaml:"dbName"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	MaxIdleCons int    `json:"maxIdleCons" yaml:"maxIdleCons"`
	MaxOpenCons int    `json:"maxOpenCons" yaml:"maxOpenCons"`
	LogMode     string `json:"logMode" yaml:"logMode"`
}
