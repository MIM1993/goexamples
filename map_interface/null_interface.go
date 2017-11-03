//interface{} 可以代表任意类型
//interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
//Go语言嵌套Map类型 http://blog.ninja911.com/blog-show-blog_id-76.html
//解析(map[string]interface{})数据格式并打印出数据 http://www.codeweblog.com/%E8%A7%A3%E6%9E%90-map-string-interface-%E6%95%B0%E6%8D%AE%E6%A0%BC%E5%BC%8F%E5%B9%B6%E6%89%93%E5%8D%B0%E5%87%BA%E6%95%B0%E6%8D%AE/
package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["int"] = 123
	m["string"] = "hello"
	m["bool"] = true

	for _, v := range m {
		switch v.(type) {
		case string:
			fmt.Println(v, "is string")
		case int:
			fmt.Println(v, "is int")
		default:
			fmt.Println(v, "is other")
		}
	}
	fmt.Println(m)

}
