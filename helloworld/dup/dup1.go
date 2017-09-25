// 代码来自《Go程序设计语言》
// 输出标准输入中出现次数大于1的行，前面是次数
// go run dup1.go
// ctrl + d退出输入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
