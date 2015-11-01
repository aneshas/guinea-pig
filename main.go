package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo("bar"))
}

func foo(arg string) string {
	return "baz"
}
