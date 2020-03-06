package extension

import (
	"fmt"
	"testing"
)

// "继承"的实现
// 不支持重写，不支持LSP，不能当成继承来用？？？（理解不深）
type Pet struct {
}

// go在有需要导出时才使用大写（在其他的包中使用）
func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) speakTo(host string) {
	p.Speak()
	fmt.Print(host)
}

type Dog struct {
	p *Pet // 子类重写后无法在取得父类的方法
	//Pet // 匿名嵌套类型,子类重写后无法在取得父类的方法
}

func (d *Dog) Speak() { // 函数的重写不会对改变父类型的函数
	fmt.Print("wang")
}

//func (d *Dog) speakTo(host string) {
//	d.Speak()
//	d.p.speakTo(host)
//}

func TestDog(t *testing.T) {
	d := new(Dog)
	d.p.Speak()
	d.Speak()
	// d.speakTO()   // 子类没有重写改方法不能使用重写后的方法
	d.p.speakTo("vin")
}
