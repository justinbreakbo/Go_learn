package string

import "testing"

func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(s)
	t.Log(len(s))
	// s[1] = "q"  // string是一个不可变的byte切片，不能给其子元素赋值
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	// go使用 rune 可以取出字符的Unicode
	s = "严"
	c := []rune(s) // 注意使用方式
	t.Log(len(c))
	t.Logf("严 uniCode is %x", c[0])
	t.Logf("严 UTF8 is %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
