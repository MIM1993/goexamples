// 代码来自《GO语言实战》
// 这个示例程序展示如何使用有缓冲通道和固定数目的goroutine来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的goroutine数量
	taskLoad         = 10 // 要处理的工作的数量
)

// wg用来等待程序完成
var wg sync.WaitGroup

// init初始化包，Go语言运行时会在其他代码执行之前优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建一个有缓冲的通道
	tasks := make(chan string, taskLoad)

	// 启动goroutine来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作处理完时关闭通道，以便所有goroutine退出
	// 当通道关闭后，goroutine依旧可以从通道接收数据，但是不能再向通道发送数据
	// 能够从已经关闭的通道接收数据这点非常重要，因为这允许通道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

// worker作为goroutine启动来处理从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已关闭
			fmt.Printf("Worker %d: Shutdown\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker %d: Started %s\n", worker, task)

		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker %d: Completed %s\n", worker, task)
	}
}
