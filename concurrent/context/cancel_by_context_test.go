package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

/*
我们常常在一些需要主动取消长时间的任务时，创建这种类型的 Context，然后把这个Context 传给长时间执行任务的 goroutine。
当需要中止任务时，我们就可以 cancel 这个Context，这样长时间执行任务的 goroutine，就可以通过检查这个 Context，知道
Context 已经被取消了。
*/
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
