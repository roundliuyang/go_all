package lb

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"log"
	"math/rand"
	"sync"
)

const WEIGHT_LOAD_BALANCE = "weight_load_balance"
const MAX_WEIGHT = 10 // 可设置的最大权重
const MIN_WEIGHT = 1  // 可设置的最小权重

// 注册自定义权重负载均衡器
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(WEIGHT_LOAD_BALANCE, &weightPikerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type weightPikerBuilder struct {
}

// 根据负载均衡策略 生成重复的连接
func (p *weightPikerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	log.Println("weightPikerBuilder build called...")

	// 没有可用的连接
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	// 此处有坑，为什么长度给0,而不是1???

	scs := make([]balancer.SubConn, 0, len(info.ReadySCs))

	for subConn, subConnInfo := range info.ReadySCs {
		v := subConnInfo.Address.BalancerAttributes.Value(WeightAttributeKey{})
		w := v.(WeightAddrInfo).Weight

		// 限制可以设置的最大最小权重，防止设置过大创建连接数太多
		if w < MIN_WEIGHT {
			w = MIN_WEIGHT
		}

		if w > MAX_WEIGHT {
			w = MAX_WEIGHT
		}

		// 根据权重 创建多个重复的连接 权重越高个数越多
		for i := 0; i < w; i++ {
			scs = append(scs, subConn)
		}

	}

	return &weightPiker{
		scs: scs,
	}
}

type weightPiker struct {
	scs []balancer.SubConn
	mu  sync.Mutex
}

// 从build方法生成的连接数中选择一个连接返回

func (p *weightPiker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {

	// 随机选择一个返回，权重越大，生成的连接个数越多，因此，被选中的概率也越大
	log.Println("weightPiker Pick called...")
	p.mu.Lock()
	index := rand.Intn(len(p.scs))
	sc := p.scs[index]
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}

type WeightAttributeKey struct{}

type WeightAddrInfo struct {
	Weight int
}
