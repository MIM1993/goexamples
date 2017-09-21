// 代码来自《GO语言实战》
// 这个示例程序展示如何使用互斥锁来定义一段需要同步访问的代码临界区资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter是所有goroutine都需要增加其值的变量
	counter int

	// wg用来等待程序结束
	wg sync.WaitGroup

	// mutex用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待goroutine结束
	wg.Wait()
	fmt.Printf("Final Counter:%d\n", counter)
}

// incCounter使用互斥锁来同步并保证安全访问
func incCounter(id int) {
	// 在函数退出时调用Done来通知main函数工作完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			// 获取counter的值
			value := counter
			// 当前goroutine从线程退出
			runtime.Gosched()
			// 增加本地value变量的值
			value++
			// 将该值保存回counter
			counter = value
		}
		// 释放锁，允许其他正在等待的goroutine进入临界区
		mutex.Unlock()
	}
}
