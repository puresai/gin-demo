package main

import (
	"sync"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i:=1;i<10;i++ {
			fmt.Println("A:",i)
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1;i<10;i++ {
			fmt.Println("B:",i)
		}
	}()

	wg.Wait()
}