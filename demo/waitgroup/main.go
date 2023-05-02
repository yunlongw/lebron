package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d start\n", n)

			// 模拟一些耗时操作
			for j := 0; j < 100000000; j++ {
				_ = j
			}

			fmt.Printf("Goroutine %d end\n", n)
		}(i)
	}

	// 等待所有的 goroutine 执行完毕
	wg.Wait()

	fmt.Println("All goroutines finished.")
}
