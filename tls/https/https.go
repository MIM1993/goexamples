//golang中使用HTTPS
//http://blog.csdn.net/wangshubo1989/article/details/77508738
//生成私钥（rsa算法）：
//openssl genrsa -out server.key 2048
//生成证书：
//openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err)
	}
}
