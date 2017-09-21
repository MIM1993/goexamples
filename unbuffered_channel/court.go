// 代码来自《GO语言实战》
// 这个示例程序展示如何用无缓存的通道来模拟2个goroutine间的网球比赛
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg用于等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓存的通道
	court := make(chan int)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("Lily", court)
	go player("Lucy", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

// player模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done来通知main函数工作完成
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道关闭，就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 通过随机数来判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 输了，关闭通道
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打给对手
		court <- ball
	}
}
