package task_cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// context 是go自带的一个取消任务的包,可以取消同时取消任务及其子任务(建议了解context)
func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
