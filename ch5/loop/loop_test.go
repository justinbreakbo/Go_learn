package loop

import "testing"

// 循环只有for一种形式
// 常用形式和c++一样

// while 循环表示方式 while(i < 5 )
func TestWhileLoop(t *testing.T) {
	i := 0
	for i < 5 {
		t.Log(i)
		i++
	}
}

// 无限循环表示方式  while(true)
// for { }
