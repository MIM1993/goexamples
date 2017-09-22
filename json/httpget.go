// 代码改编自《GO语言实战》
// 并参考自https://studygolang.com/articles/4335
// 处理Get请求响应的JSON示例
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type gResponse struct {
	GcardNumber   uint64 `json:"card_number"`
	GcardBalance  string `json:"card_balance"`
	GbalanceTime  string `json:"balance_time"`
	GcardValidity string `json:"card_validity"`
	GcurrentTime  string `json:"current_time"`
}

func main() {
	uri := "http://api.oupag.com/dev/api/shenzhentong.php?cardno=29444"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	// 将JSON响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// %+v打印结构体时，会添加字段名Printf("%+v", people)  {Name:zhangsan}
	// https://studygolang.com/articles/2644
	fmt.Printf("%+v", gr)
}
