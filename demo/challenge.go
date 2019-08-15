/**
goroutine资源竞争，锁的使用
统一资源的读写必须是原子性的
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg sync.WaitGroup
	mt sync.Mutex
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()

	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i:=0;i<5;i++ {
		mt.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mt.Unlock()
	}
}