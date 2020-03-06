package client

import (
	"Go_learn/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	fib := series.GetFibonacci(10)
	t.Log(fib)
}

// go无法在同一环境下，不同项目使用同一个包的不同版本；无法管理对包的特定版本的依赖‘
//所以go1.5版本之后加入了vendor路径，用来解决这个问题
//可以用一些常用的依赖（包）管理工具:godep/glide/dep (在github上面找)
