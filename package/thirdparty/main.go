// 代码来自https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.5.md
// 测试第三方包
// 导入第三方包：go get github.com/fengchunjian/goexamples/package/userdefined/pack1
// 编译方式：go build main.go
package main

import (
	"fmt"
	"github.com/fengchunjian/goexamples/package/userdefined/pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
}
