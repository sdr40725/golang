package main

import "fmt"

const (
	test001 = iota + 1
	test002
)

func main() {

	foo := "Hello"
	bar := 100
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(test001)
	fmt.Println(test002)
}
