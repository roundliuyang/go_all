package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/**
使用通过复用，降低复杂对象的创建和GC代价
协程安全，会有锁的开销
生命周期受GC的影响，不适合做连接池等，需要自己管理生命周期的资源的池化
*/
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	fmt.Println("--")
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	time.Sleep(1 * time.Second)
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
