package main

import "fmt"

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	x += 1
	fmt.Println(x)
}
