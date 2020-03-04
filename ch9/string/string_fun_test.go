package string

import (
	"strconv"
	"strings"
	"testing"
)

// 常用字符串函数
//strings包 https://golang.org/pkg/strings/
//strconv包 https://golang.org/pkg/strconv/

func TestStringFun(t *testing.T) {
	// 分割字符串   字符串->数组
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}

	// 链接字符串 数组->字符串
	arr := []string{"1", "2", "3"}
	t.Log(strings.Join(arr, "-"))

	// 字符串和其他类型的转换（尤其是整数）
	s = strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
