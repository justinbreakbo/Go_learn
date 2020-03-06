package customType

import (
	"fmt"
	"testing"
	"time"
)

// 自定义类型
type IntConv func(op int) int

// 自定义一个函数类型 简化函数写法
func timeSpent(inner IntConv) IntConv { // inner为传入的函数的函数名
	return func(n int) int {
		start := time.Now()
		ret := inner(n) // 返回的函数执行该入参函数，该入参函数的参数为n，在返回函数执行时传入
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
