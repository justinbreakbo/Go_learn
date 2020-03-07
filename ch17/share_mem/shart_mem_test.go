package share_mem

import (
	"sync"
	"testing"
	"time"
)

//传统的共享内存的并发机制

func TestShareMemory(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++ // 并发竞争条件里面做自增（线程不安全的程序）
		}()
	}
	//time.Sleep(1 * time.Second)
	t.Log("counter = ", counter)
}

//需要对共享的内存做锁的保护，以保证线程安全
func TestShareMemoryThreadSafe(t *testing.T) {
	// 此处使用了Mutex，它只有一种锁；平时建议使用rwLock，它分为读锁和写锁，效率会更高(读锁和读锁之间不是互斥，不会阻塞)
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock() // 在defer里面解锁，防止最后忘记解锁导致阻塞
			}()
			mut.Lock()
			counter++
		}()
	}
	// 外面的程序执行得太快，里面协程还没执行完就已经退出，会导致得不到正确的结果，需要用Sleep避免
	time.Sleep(1 * time.Second)
	t.Log("counter = ", counter)
}

// 使用WaitGroup可以准确的等待，不会浪费太多等待的时间，优于Sleep方法
//（一个同步各个线程的方法）
func TestShareMemoryWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 增加等待的量
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 任务结束主动告诉程序
		}()
	}
	wg.Wait() // 最后等待协程执行完
	t.Log("counter = ", counter)
}
