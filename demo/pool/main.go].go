package main

import (
	"fmt"
	"sync"
)

// 定义任务类型
type Task func()

// 协程池类型
type Pool struct {
	wg    sync.WaitGroup // 用于等待所有任务执行完成
	size  int            // 协程池大小
	tasks chan Task      // 任务队列
}

// 创建一个新的协程池
func NewPool(size int) *Pool {
	return &Pool{
		size:  size,
		tasks: make(chan Task),
	}
}

// 启动协程池，开始执行任务
func (p *Pool) Start() {
	for i := 0; i < p.size; i++ {
		go p.worker()
	}
}

// 关闭协程池，等待所有任务执行完成
func (p *Pool) Stop() {
	close(p.tasks)
	p.wg.Wait()
}

// 将任务添加到任务队列中
func (p *Pool) AddTask(task Task) {
	p.tasks <- task
}

// 工作协程，用于执行任务
func (p *Pool) worker() {
	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func main() {
	// 创建一个大小为 3 的协程池
	pool := NewPool(3)

	// 添加 5 个任务到任务队列中
	for i := 1; i <= 5; i++ {
		taskID := i
		pool.wg.Add(1)
		pool.AddTask(func() {
			fmt.Printf("Task %d is running...\n", taskID)
		})
	}

	// 启动协程池
	pool.Start()

	// 等待所有任务执行完成
	pool.Stop()

	fmt.Println("All tasks have been completed.")
}
