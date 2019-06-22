/**
单向通道
*/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

// 只能发送
func producer(send chan<- int) {
	for i:=0;i<4;i++ {
		send <- i
		// i <- send
	} 
	close(send)
}

//只能接收
func consumer(receive <-chan int) {
	for num:= range receive {
		fmt.Println("value=", num)
	}
}