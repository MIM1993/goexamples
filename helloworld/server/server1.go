// 代码来自《Go程序设计语言》P15
// server1是一个迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 回声请求调用处理程序
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// 处理程序回显请求
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
