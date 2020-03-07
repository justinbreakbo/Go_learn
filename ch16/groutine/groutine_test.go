package groutine

import (
	"fmt"
	"testing"
	"time"
)

//协程（轻量级）与系统线程是多对多的关系

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 60) // Test执行过快，加了这句才能输出
}
