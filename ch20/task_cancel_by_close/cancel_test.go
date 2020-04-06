package task_cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 不安全
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // struct{} 表示空结构,第二个{}表示实例化这个空结构
}

// 用close方法更安全
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "canceled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
