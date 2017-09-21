// 代码改编自https://studygolang.com/articles/6621
// runtime.Gosched()用于让出CPU时间片
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg用来等待程序结束
var wg sync.WaitGroup

// say打印两次string s
func say(s string) {
	// 在函数退出时调用Done来通知main函数工作完成
	defer wg.Done()

	for i := 0; i < 2; i++ {
		fmt.Println(s)
		runtime.Gosched()
		// 如果注释掉runtime.Gosched()将连续打印两次s
	}
}

func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go say("hello")
	go say("world")

	// 等待goroutine结束
	wg.Wait()
}
