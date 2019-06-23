/**
读写锁的使用 sync.RWMutex
*/
package main

import (
	"fmt"
	"sync"
	"math/rand"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func main() {
	wg.Add(10)

	for i:=0; i < 5;i++ {
		go read(i)
	}

	for i:= 0;i<5;i++ {
		go write(i)
	}

	wg.Wait()
}


func read(n int) {
	rw.RLock()
	fmt.Println("读取操作", n)

	v := count
	fmt.Println("读取结束", n, v)

	wg.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Println("写入操作")

	v := rand.Intn(100)

	count = v
	fmt.Println("写入结束")
	wg.Done()
	rw.Unlock()
}