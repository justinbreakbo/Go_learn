package error

import (
	"errors"
	"fmt"
	"testing"
)

//go没有异常机制 （没有try—catch）
//error类型实现了error接口
//可以通过error.New来快速创建错误实例

//可以定义不同的错误变量，以便于判断错误类型
var LessThanTwoError error = errors.New("n must be greater than 2")
var GreaterThanHundredError error = errors.New("n must be less than 100")

func getFibonacci(n int) ([]int, error) {
	// 及早失败，避免嵌套
	//if n < 2 || n > 100 {
	//	return nil, errors.New("n should be in [2,100]")
	//}
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, GreaterThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	// 及早失败，避免嵌套
	if v, err := getFibonacci(11); err == LessThanTwoError {
		fmt.Println("it is less.")
		t.Error(err)
	} else if err == GreaterThanHundredError {
		fmt.Println("it is greater.")
		t.Error(err)
	} else {
		t.Log(v)
	}
}
