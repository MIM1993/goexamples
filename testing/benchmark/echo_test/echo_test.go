// 代码为《Go程序设计语言练习1.3
// echo程序性能测试
package echo_test

import (
	//"fmt"
	"os"
	"strings"
	"testing"
)

// go test -v -run="none" -bench="BenchmarkEcho1" -benchtime="3s"
func BenchmarkEcho1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 1; j < len(os.Args); j++ {
			s += sep + os.Args[j]
			sep = " "
		}
		//fmt.Println(s)
	}
}

// go test -v -run="none" -bench="BenchmarkEcho2" -benchtime="3s"
func BenchmarkEcho2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		//fmt.Println(s)
	}
}

// go test -v -run="none" -bench="BenchmarkEcho3" -benchtime="3s"
func BenchmarkEcho3(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
		//fmt.Println(strings.Join(os.Args[1:], " "))
	}
}
