package _interface

import (
	"fmt"
	"testing"
)

// 接口，定义对象之间交互协议

type Programmer interface {
	writeHelloWorld() string
}

// go语言不用implement或extend之类的关键字来声明对接口的依赖（非入侵性），因此不会发生循环依赖（接口可以定义在自己的包里面）
type GoProgrammer struct {
	name string
}

// duckType 鸭子类型 方法签名一样就是实现该接口(这种设置可以先定义方法，后来发现很多类似的方法时，再写接口）
func (g GoProgrammer) writeHelloWorld() string {
	return fmt.Sprintf("hello world,%s", g.name)
}

func TestInterface(t *testing.T) {
	var p Programmer // 需要先实例接口（定义接口变量）
	p = new(GoProgrammer)
	t.Log(p.writeHelloWorld())
}

//var prog coder = &GoProgrammer{}
//接口变量包括类型和数据
//类型： type GoProgrammer struct{}
//数据： &Goprogrammer{}
