package series

import "fmt"

// 在main被执行前，所依赖的package的init函数都会被执行，执行的顺序由包导入的依赖关系决定
// 每个包可以有多个init函数，包的每个源文件也可以有多个init函数，且它们都会被执行（奇特！）
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

// 首字母大写 才能在包外被访问
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList
}
