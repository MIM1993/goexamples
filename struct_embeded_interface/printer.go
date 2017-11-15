//在结构体（struct）中内嵌 接口（interface）
//https://studygolang.com/articles/6934
package main

import "fmt"

type Printer interface {
	Print()
}

type CanonPrinter struct {
	printerName string
}

func (printer CanonPrinter) Print() {
	fmt.Println(printer.printerName, "print.")
}

type PrintWoker struct {
	Printer
	name string
	age  int
}

func main() {
	canon := CanonPrinter{printerName: "canon_1"}
	printWorker := PrintWoker{Printer: canon, name: "Zhang", age: 21}
	printWorker.Printer.Print()
}
