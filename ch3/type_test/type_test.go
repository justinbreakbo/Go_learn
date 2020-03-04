package type_test

import "testing"


// 不支持隐式类型转换
func TestImplicit(t *testing.T)  {
	var a int16 = 1
	var b int
	// b = a	// 不支持
	b = int (a)
	t.Log(a,b)
}

// 别名的隐式类型转换也不支持
type MyInt int64
func TestAlias(t *testing.T)  {
	var a int16 = 1
	var b MyInt
	// b = a   // 不支持
	b = MyInt(a)
	t.Log(a,b)
}

func TestPoint(t *testing.T)  {
	var a = 1
	var aPtr = &a
	//aPtr = aPtr + 1		// 不能进行指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}


func TestString(t *testing.T) {
	var a string	//默认初始化字符串为空字符串，而不是空
	t.Log("*"+ a +"*")
	t.Log(len(a))
	if a == "" {
		t.Log("a is a nil string")
	}
}



