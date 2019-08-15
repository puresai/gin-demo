/**
通道，在两个goroutine之间架设的管道，可以进行数据交互
无缓冲通道(大小为0)
*/
package main

import (
	"fmt"
)

func main() {
	one := make(chan int)
	two := make(chan int)

	go func() {
		one<-100
	}()

	go func() {
		v:=<-one
		two<-v
	}()

	fmt.Println(<-two)
}