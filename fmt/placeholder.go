// 代码改编自https://studygolang.com/articles/2644
// golang fmt格式占位符
package main

import "fmt"

type Human struct {
	Name string
}

var people = Human{Name: "zhangsan"}

func main() {
	// %v相应值的默认格式，Printf("%v", people)输出{zhangsan}
	fmt.Printf("%v\n", people)
	// %+v打印结构体时，会添加字段名，Printf("%+v", people)输出{Name:zhangsan}
	fmt.Printf("%+v\n", people)
	// %#v相应值的Go语法表示，Printf("%#v", people)输出main.Human{Name:"zhangsan"}
	fmt.Printf("%#v\n", people)
	// %T相应值的类型的Go语法表示，Printf("%T", people)输出main.Human
	fmt.Printf("%T\n", people)
	fmt.Println()

	// 字面上的百分号，并非值的占位符
	fmt.Printf("%%\n")
	// 布尔占位符true或false
	fmt.Printf("%t\n", true)
	// 指针占位符
	fmt.Printf("%p\n", &people)
	fmt.Println()

	// 二进制表示
	fmt.Printf("%b\n", 5)
	// 相应Unicode码点所表示的字符
	fmt.Printf("%c\n", 0x4E2D)
	// 十进制表示
	fmt.Printf("%d\n", 0x12)
	// 八进制表示
	fmt.Printf("%d\n", 10)
	// 单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("%q\n", 0x4E2D)
	// 十六进制表示，字母形式为小写 a-f
	fmt.Printf("%x\n", 13)
	// 十六进制表示，字母形式为大写 A-F
	fmt.Printf("%X\n", 13)
	// Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("%U\n", 0x4E2D)
	fmt.Println()

	// 科学计数法，例如 -1234.456e+78
	fmt.Printf("%e\n", 10.2)
	// 科学计数法，例如 -1234.456E+78
	fmt.Printf("%E\n", 10.2)
	// 有小数点而无指数，例如 123.456
	fmt.Printf("%f\n", 10.2)
	fmt.Println()

	// 输出字符串表示（string类型或[]byte)
	fmt.Printf("%s\n", []byte("Go语言"))
	// 双引号围绕的字符串，由Go语法安全地转义
	fmt.Printf("%q\n", "Go语言")
	// 十六进制，小写字母，每字节两个字符
	fmt.Printf("%x\n", "golang")
	// 十六进制，大写字母，每字节两个字符
	fmt.Printf("%X\n", "golang")
}
