package encapsulation

import (
	"fmt"
	"testing"
)

// go没有面向对象，因为它没有继承，只有封装和多态

// 封装的实现

// 结构体定义
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 方法定义
// 第一种方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) string1() string {
	return fmt.Sprintf("id:%s-name:%s-age:%d", e.Id, e.Name, e.Age)
}

// 第二种方法不会进行值复制，所以为了避免内存拷贝，减少开销，建议使用第二种定义方式
func (e *Employee) string2() string {
	return fmt.Sprintf("id:%s/name:%s/age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{"0", "Bob", 30}
	e1 := Employee{Name: "Jim", Age: 20}
	e2 := new(Employee) // 这种定义方式会返回对象指针,但访问其成员或方法时都不需要用"->"符号
	e2.Name = "Marry"
	e2.Age = 22
	e2.Id = "2"
	t.Log(e)
	t.Log(e.Age)
	t.Log(e1)
	t.Log(e1.Name)
	if e1.Id == "" {
		t.Log("e1.id is null")
	} // 未初始化的值一般都是零值或空字符串
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)

	t.Log(e.string1())
	t.Log(e.string2())
}
