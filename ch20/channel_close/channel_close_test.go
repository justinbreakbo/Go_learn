package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 通道关闭(出现一种广播的机制)
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 数据传送完之后关闭channel,关闭后不能再往channel发消息,否则会产生panic
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i := 0; i < 11; i++ { // 现实情况下应该是不知道生产者有多少个数据
		//	data := <-ch			// 通道关闭后,在该通道取消息会返回一个通道类型的零值
		//	fmt.Println(data)
		//}
		for {
			// ok为bool值,True表示正常接收,False表示通道关闭
			// (这种广播机制很常用,多个receiver可以根据这个信息执行退出之类的操作)
			if data, ok := <-ch; ok {
				fmt.Println(data)

			} else { // receiver要对通道关闭状态做出处理
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1) // 两个receiver,输出的顺序可能有变化
	dataReceiver(ch, &wg)

	wg.Wait()
}
