package empty_interface

import (
	"fmt"
	"testing"
)

//空接口，可以表示任何类型，类似Object类

//通过断言可以将空接口转换成指定类型
//v, ok = p.(int)   //ok 为true 表示转换成功

func doSomething(p interface{}) {
	//if v,ok := p.(int) ; ok {
	//	fmt.Println("Integer", v)
	//}
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Print("unknown type")
	}
}

func TestEmptyInterface(t *testing.T) {
	a := 10
	b := "10"
	doSomething(a)
	doSomething(b)
}

// 接口的最佳实践
//1.倾向于使用更小的接口
//2.较大的接口应该由较小的接口组合而成
//3.依赖接口时，只依赖于必要功能的最小接口
