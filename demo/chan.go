/**
chan+select
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("exit...")
			default:
				fmt.Println("go...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10*time.Second)
	fmt.Println("ok")
	stop <- true
	time.Sleep(5*time.Second)
}