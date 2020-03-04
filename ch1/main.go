package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Args)
	// go 中获取命令行输入需要用os.Args(数组)
	if len(os.Args) > 1 {
		fmt.Println("hello world",os.Args[1])
	}

	// go 的main函数没有返回值，只能通过os.Exit()退出
	os.Exit(0)
}