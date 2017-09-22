// 参考自http://www.cnblogs.com/zhepama/archive/2013/04/12/3017230.html
// 测试time.Unix()和time.UnixNano()
// time.Unix()：自从1970年1月1号到现在的时间戳（秒）
// time.UnixNano()：自从1970年1月1号到现在的时间戳（纳秒）
// 1秒(s)=1000000000纳秒(ns)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}
