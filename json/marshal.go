// 代码来自《GO语言实战》
// 使用json包的MarshalIndent函数进行编码
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个保存键值对的map
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "Programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将这个映射序列化到JSON字符串
	// MarshalIndent类似Marshal但会使用缩进将输出格式化
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}
