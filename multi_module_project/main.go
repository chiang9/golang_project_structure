package main

import (
	"fmt"
	"multi_mod_project/cmd/mod2"
	"multi_mod_project/cmd/mod3"
	"multi_mod_project/cmd/mod3/mod3_sub"
	"multi_mod_project/mod1"
)

func main() {
	fmt.Println("Hello, world")
	mod1.Print_string("Hello, world")
	mod2.Print_string("Hello, world")
	mod3.Print_string("Hello, world")
	mod3_sub.Print_string("Hello, world")
}
