package condition

import "testing"

// if 条件语句支持两段，第一段赋值，第二段比较
func TestIf(t *testing.T)  {
	if a := 1 == 1; a {
		t.Log(a)
		t.Log("1=1")
	}
}

// switch 语句不局限于常量或整型，字符串也可以
// go里面switch不需要加break（默认已经加了，命中一个case后会自动中断）
// case 里面可以放多个结果选项，用逗号分割（命中一个即可）
// 可以不设定switch语句，此时整个switch结构就和多个 if...else if... 一样
func TestSwitchMutiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0,2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T){
	for i := 0; i < 5; i++ {
		switch  {
		case i%2 == 0 :
			t.Log("even")
		case i%2 != 0:
			t.Log("odd")
		default:
			t.Log("it is err")
		}
	}
}
