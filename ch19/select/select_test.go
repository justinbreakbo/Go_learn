package _select

import (
	"fmt"
	"testing"
	"time"
)

//多路选择(select)实现超时控制

func service() string {
	time.Sleep(time.Millisecond * 200)
	return "done"
}

func asyncService() chan string {
	retCh := make(chan string) // 不带buffer的channel
	//retCh := make(chan string, 1) // 带buffer的channel,通常更高效
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
