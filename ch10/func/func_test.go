package _func

import (
	"fmt"
	"testing"
	"time"
)

// go中所有参数都是值传递，slice、map、channel会有传引用的错觉

// 函数返回多个值
func returnMultiValues(op int) (int, int) {
	return op, op * op
}

// 函数的参数，返回值都可以是函数
// 函数式编程（较复杂）
func timeSpent(inner func(op int) int) func(op int) int { // inner为传入的函数的函数名
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
	a, b := returnMultiValues(3)
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 可变参数（实际上不管传入多少个参数，最终都转变成一个数组来遍历完成）
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5))
	t.Log(sum(1, 2, 3, 4))
}

// 延迟执行 (在返回之前才执行，通常用于释放资源，释放锁等）
func clear() {
	fmt.Println("Clear resourses.")
}

func TestDefer(t *testing.T) {
	defer clear() //会在最后执行（return之前）
	fmt.Println("Start.")
	panic("err") // defer 的语句在panic之后依然会执行
	//fmt.Println("repaired")  //其他的语句在panic语句后不会执行
}
