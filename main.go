package main

import (
	"fmt"
	"github.com/sdr40725/golang/helloworld"
)

func main() {
	a := 1
	fmt.Printf(helloworld.HelloWorld())
	if a > 1 {
		fmt.Println("a > 1")
	}
}
