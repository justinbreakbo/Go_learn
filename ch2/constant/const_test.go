package constant

import "testing"

// 连续常量定义的简便方法，iota为自增，遇到新的const会重新置零
const (
	Monday = 1 << iota
	Tuesday
	Wednesday
)

func TestConstant(t *testing.T)  {
	a := 3
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Monday&a == Monday, Tuesday&a == Tuesday, a&Wednesday == Wednesday)
	
}