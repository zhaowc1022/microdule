package rpc

import (
	grpcs "github.com/hihibug/microdule/v2/rpc/grpc"
	etcdClientV3 "go.etcd.io/etcd/client/v3"
)

type Rpc interface {
	Client() any
	Register(*etcdClientV3.Client) (*grpcs.ServiceRegister, error)
	Run() error
	Close()
}
