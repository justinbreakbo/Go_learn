package csp

import (
	"fmt"
	"testing"
	"time"
)

// go语言独有的并发机制:csp
//channel有两种:1.通信两方同时在channel上; 2.带buffer(有容量)的channel

// 这是一个串行的结构
func service() string {
	time.Sleep(time.Millisecond * 100)
	return "done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// 改造成异步结构(使用channel)
func asyncService() chan string {
	retCh := make(chan string) // 不带buffer的channel
	//retCh := make(chan string, 1) // 带buffer的channel,通常更高效
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited") // 带buffer时不用等待取走channel的东西就可以执行下一步
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
