package array

import "testing"

//数组不可伸缩
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [2][2]int{{1, 2}, {3, 4}} // 二维数组
	arr3 := [...]int{1, 3, 4}         // 不写数组长度可用...代替，自动识别长度
	arr4 := []int{1, 2, 2}            // 直接不写长度是声明切片slice，不是数组，应注意
	//arr = append(arr, 1)				// 数组不可变长，在其声明时就已经确定长度

	t.Log(arr, arr1, arr2, arr3, arr4)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	// 使用range关键字协助遍历
	for id, value := range arr {
		t.Log(id, value)
	}
}

// 数组截取	（切片可以用数组截取的方式来获得）
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec := arr[:3]
	//arr_sec1 := arr [:-1]    	// go不支持负数的切片,python支持
	t.Log(arr_sec)

}
