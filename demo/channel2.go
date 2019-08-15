/**
有缓冲通道(类似队列，发送插入到尾部，满了会阻塞；
接收从头部取出并删除，空了会阻塞，可以通过cap和len函数获取通道最大容量和现有元素数量)
单向通道
*/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	fmt.Println(cap(ch))
	fmt.Println(len(ch))

	
}