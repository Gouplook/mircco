package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"koalamicrro/registry"
)

//etcd 注册插件
type EtcdRegistry struct {

	options *registry.Options
	client *clientv3.Client
	serviceCh chan * registry.Service
}

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

type RegisterServie struct {
	id clientv3.LeaseID
	service *registry.Service
	registered bool
	keeAliveCh <-chan *clientv3.LeaseKeepAliveResponse

}

//插件的名字

func (e *EtcdRegistry)Name()string{
	return "etcd"

}
func (e *EtcdRegistry)Init(ctx context.Context,opts ...registry.Option)(err error){

	return
}

//服务注册
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {
	return
}
//服务反注册
func (e *EtcdRegistry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	return
}