package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	numRequests := 1000000 // 想要模拟的请求数
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080/ping") // 替换为实际的服务器地址和端口
			if err != nil {
				fmt.Printf("Error making the request: %v\n", err)
				return
			}
			defer resp.Body.Close()

			// 读取响应并处理
			if resp.StatusCode == http.StatusOK {
				fmt.Println("Request successful")
			} else {
				fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
			}
		}()
		time.Sleep(2 * time.Millisecond) // 为了模拟请求之间的延迟
	}

	wg.Wait() // 等待所有请求完成
	fmt.Println("All requests completed.")
}
