package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

// panic 用于不可恢复的错误，panic退出前会执行defer指定的内容(注意是退出前执行defer，defer还是最后执行的）

// os.Exit 退出前不会调用defer指定的内容、不输出当前调用栈信息

func TestPanicVsExit(t *testing.T) {
	defer func() {
		// recover() 返回的错误就是panic传进去的错误
		if err := recover(); err != nil {
			// 此处应根据不同的err执行处理，而不是直接把错误记录下来而已
			// Let it crash! 是恢复不确定性错误的最好方法（health check 会把系统重启）
			fmt.Println("recover from ", err)
		}
	}()
	fmt.Println("start")

	panic(errors.New("something wrong"))
	//os.Exit(0)
}
