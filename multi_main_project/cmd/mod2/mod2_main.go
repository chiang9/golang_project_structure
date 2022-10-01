package main

import (
	"fmt"

	"multi_main_project/cmd/func1"
	"multi_main_project/cmd/mod1/mod1_func"
)

func main() {
	fmt.Println("Hello, world from mod2")
	func1.Print_string("Hello, world")
	mod1_func.Print_string("Hello, world")
}
