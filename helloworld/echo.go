// 代码来自《GO程序设计语言》
// 输出命令行参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// echo1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// echo2
	fmt.Println(s)

	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// echo3
	fmt.Println(strings.Join(os.Args[1:], " "))

	// 练习1.1
	fmt.Println(os.Args[0])

	// 练习1.2
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
