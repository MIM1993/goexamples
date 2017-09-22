// 代码来自《GO语言实战》
// 这个示例程序展示如何使用pool包来共享一组模拟的数据库连接
package main

import (
	"./pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 5 // 要使用的goroutine数量
	pooledResources = 2 // 池中的资源的数量
)

// dbConnection模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close实现了io.Closer接口，以便dbConnection可以被池管理，Close用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter用来给每个连接分配一个独一无二的ID
var idCounter int32

// createConnection是一个工厂函数，当需要一个新连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 1; query <= maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 等待goroutine结束
	wg.Wait()

	// 关闭池
	log.Println("Shutdown Program.")
	p.Close()
}

// performQueries用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将这个连接释放回池里
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// conn.(*dbConnection)为接口类型断言，参考https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.3.md
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
