// 代码来自《GO语言实战》
// 这个示例程序实现了简单的网络服务
package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener Started, Listening on:4000")
	http.ListenAndServe(":4000", nil)
}
