// 代码来自《GO语言实战》
// 这个示例展示来自不同标准库的不同函数是如何使用io.Writer接口的
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer值，并将这个字符串写入Buffer
	// 使用实现io.Writer的Write方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用Fprintf将一个字符串拼接到Buffer中
	// 将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&b, "World!")

	// 将Buffer内容输出到标准输出设备
	// 将os.File值得地址，作为io.Writer类型值传入
	b.WriteTo(os.Stdout)
}
