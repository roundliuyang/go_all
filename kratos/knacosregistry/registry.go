package knacosregistry

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var ErrServiceInstanceNameEmpty = errors.New("kratos/knacos: ServiceInstance.Name can not be empty")

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type Option = func(r *Registry)

func WitchCluster(cluster string) Option {
	return func(r *Registry) {
		r.Cluster = cluster
	}
}

func WithGroup(group string) Option {
	return func(r *Registry) {
		r.Group = group
	}
}

func WithKind(kind string) Option {
	return func(r *Registry) {
		r.Kind = kind
	}
}

func WithWeight(w float64) Option {
	return func(r *Registry) {
		r.Weight = w
	}
}

type Registry struct {
	Cluster string
	Group   string
	Kind    string
	Weight  float64

	nc naming_client.INamingClient
}

func NewRegistry(nc naming_client.INamingClient, opts ...Option) *Registry {
	r := &Registry{
		Cluster: "DEFAULT",
		Group:   constant.DEFAULT_GROUP,
		Kind:    "grpc",
		Weight:  100,
		nc:      nc,
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

// Register the registration.
func (r *Registry) Register(ctx context.Context, si *registry.ServiceInstance) error {
	if si.Name == "" {
		return ErrServiceInstanceNameEmpty
	}

	for _, endpoint := range si.Endpoints {
		u, e := url.Parse(endpoint)
		if e != nil {
			return e
		}
		ip, port, e := net.SplitHostPort(u.Host)
		if e != nil {
			return e
		}
		p, e := strconv.ParseUint(port, 10, 64)
		if e != nil {
			return e
		}

		var md map[string]string
		if si.Metadata == nil {
			md = map[string]string{
				"kind":    u.Scheme,
				"version": si.Version,
			}
		} else {
			md = make(map[string]string, len(si.Metadata)+2)
			for k, v := range si.Metadata {
				md[k] = v
			}
			md["kind"] = u.Scheme
			md["version"] = si.Version
		}

		if _, e = r.nc.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          ip,
			Port:        p,
			Weight:      r.Weight,
			Enable:      true,
			Healthy:     true,
			Metadata:    md,
			ClusterName: r.Cluster,
			ServiceName: si.Name,
			GroupName:   r.Group,
			Ephemeral:   true,
		}); e != nil {
			return fmt.Errorf("RegisterInstance %v err: %w", endpoint, e)
		}
	}

	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	for _, endpoint := range service.Endpoints {
		u, e := url.Parse(endpoint)
		if e != nil {
			return e
		}
		host, port, e := net.SplitHostPort(u.Host)
		if e != nil {
			return e
		}
		p, e := strconv.ParseUint(port, 10, 64)
		if e != nil {
			return e
		}
		if _, e = r.nc.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          host,
			Port:        uint64(p),
			ServiceName: service.Name,
			GroupName:   r.Group,
			Cluster:     r.Cluster,
			Ephemeral:   true,
		}); e != nil {
			return fmt.Errorf("DeregisterInstance %v err: %w", endpoint, e)
		}
	}
	return nil
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	rsp, e := r.nc.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   r.Group,
		HealthyOnly: true,
	})
	if e != nil {
		return nil, e
	}

	ret := make([]*registry.ServiceInstance, len(rsp))
	for i, ins := range rsp {
		kind := r.Kind
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

// Watch creates a watcher according to the service name.
func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	w := newWatcher(ctx, r.nc, serviceName, r.Group, r.Kind, []string{r.Cluster})
	return w, w.start()
}
