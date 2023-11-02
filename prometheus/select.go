package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for {
		select {
		case <-ch:
			// 模拟处理
			time.Sleep(1 * time.Second)
			fmt.Println("ch")
		default:
		}
	}
}
