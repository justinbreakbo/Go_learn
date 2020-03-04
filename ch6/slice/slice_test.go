package slice

import (
	"testing"
)

// slice可以伸缩
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 建议用make函数声明一个slice
	s2 := make([]int, 1, 5) // 可以声明其len和cap
	t.Log(len(s2), cap(s2))
	// t.Log(s2[0], s2[1])		// 不能访问超出length 的值
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1])
}

// 切片实现可变长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // cap的增长速率是2的倍数
	}
}

// 切片共享存储结构
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown" // 后端共享存储空间，其中一个slice更改了其中的值，所有的slice都会更改
	t.Log(Q2)
	t.Log(year)
}

// 数组可以比较，切片不可比较
func TestCompare(t *testing.T) {
	a := [...]int{1, 2}
	b := [...]int{1, 2}
	if a == b {
		t.Log("equal")
	}
	//slice can only be compared to nil
	//c := []int{1, 2}
	//d := []int{1, 2}
	//if c == d {
	//	t.Log("equal")
	//}
}
