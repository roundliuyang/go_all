package knacosregistry

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	serviceName string
	clusters    []string
	groupName   string
	ctx         context.Context
	cancel      context.CancelFunc
	watchChan   chan struct{}
	nc          naming_client.INamingClient
	kind        string
}

func newWatcher(ctx context.Context, nc naming_client.INamingClient, serviceName, groupName, kind string, clusters []string) *watcher {
	subctx, cancel := context.WithCancel(ctx)
	return &watcher{
		ctx: subctx, cancel: cancel,
		serviceName: serviceName,
		clusters:    clusters,
		groupName:   groupName,
		nc:          nc,
		kind:        kind,
		watchChan:   make(chan struct{}, 1),
	}
}

func (w *watcher) start() error {
	return w.nc.Subscribe(&vo.SubscribeParam{
		ServiceName: w.serviceName,
		Clusters:    w.clusters,
		GroupName:   w.groupName,
		SubscribeCallback: func(services []model.Instance, err error) {

		},
	})
}

// Next returns services in the following two cases:
// 1.the first time to watch and the service instance list is not empty.
// 2.any service instance changes found.
// if the above two conditions are not met, it will block until context deadline exceeded or canceled
func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case <-w.watchChan:
	}

	rsp, e := w.nc.GetService(vo.GetServiceParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
	if e != nil {
		return nil, e
	}

	ret := make([]*registry.ServiceInstance, len(rsp.Hosts))
	for i, ins := range rsp.Hosts {
		kind := w.kind
		if k, ok := ins.Metadata["kind"]; ok {
			kind = k
		}

		ret[i] = &registry.ServiceInstance{
			ID:        ins.InstanceId,
			Name:      ins.ServiceName,
			Version:   ins.Metadata["version"],
			Metadata:  ins.Metadata,
			Endpoints: []string{fmt.Sprintf("%s://%s:%d", kind, ins.Ip, ins.Port)},
		}
	}
	return ret, nil
}

// Stop close the watcher.
func (w *watcher) Stop() error {
	w.cancel()

	return w.nc.Unsubscribe(&vo.SubscribeParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
}
