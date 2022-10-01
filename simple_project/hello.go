package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("Hello, world")
	Print_string("Hello, world")
	fmt.Println(stringutil.Reverse("Hello, world from 3rd party package"))
}
