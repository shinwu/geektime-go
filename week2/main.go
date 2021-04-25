package main

import (
	"fmt"
	"geektime-go/wrap"
)

func main() {
	fmt.Println("使用 go 内制的 error 处理")
	wrap.Way1()

	fmt.Println()

	fmt.Println("使用 github.com/pkg/errors 处理")
	wrap.Way2()
}
