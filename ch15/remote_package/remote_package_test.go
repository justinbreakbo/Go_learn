package remote_package

// 远程的包用go get命令获取
// go get -u 强制从网络上更新远程依赖
// eg: go get -u github.com/easierway/concurrent_map
// 把自己写的包上传到github，直接从代码路径开始，不要有src

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
