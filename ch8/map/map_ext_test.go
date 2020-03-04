package map_ext

import "testing"

// map的value可以是一个方法(可以用来实现工厂模式）
func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// go里面没有set，可以用map实现set（排好序并且不重复） map[int]bool
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	t.Log(mySet)
	n := 2
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
