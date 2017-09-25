// 代码来自《GO语言实战》
// handlers包提供了用于网络服务的服务端点
package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON返回一个简单的JSON文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@email.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
