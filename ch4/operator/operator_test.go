package operator

import "testing"

//运算符：算术、比较、逻辑、位

// 数组可以比较是否相等，但需要两个数组的元素个数一致

// 按位置零运算符 &^   (Go 特有)
func TestBitClear(t *testing.T)  {
	a := 1
	t.Log(a)
	a = a &^0	// 右边的二进制位为1，则清零；右边的二进制位为0，则保留左边原来数值
	t.Log(a)
}



