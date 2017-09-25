// 代码来自《GO语言实战》
// 这个示例程序展示如何使用io.Reader和io.Writer接口写一个简单版本的curl
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 这里的r是一个响应，r.Body是io.Reader
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// 创建文件保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// 使用MultiWriter，可以同时向文件和标准输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	// 读出响应的内容，并写到两个目的地
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}
