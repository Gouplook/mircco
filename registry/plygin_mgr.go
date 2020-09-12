package registry

import (
	"context"
	"sync"
)

var (
	pluginmgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

type PluginMgr struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

func (p *PluginMgr) registryPlugin(plugin Registry) (err error) {

	return
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string,
	opts ...Option) (registry Registry, err error) {

	return
}

// 注册插件
func RegisterPlugin(registry Registry) (err error) {
	return pluginmgr.registryPlugin(registry)
}

//初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginmgr.initRegistry(ctx, name, opts...)
}
