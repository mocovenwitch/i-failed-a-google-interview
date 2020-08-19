package main

import "fmt"

func main() {
	var a string
	fmt.Println(len(a))
	fmt.Println(&a)

	a = "hello"
	fmt.Println(len(a))
	fmt.Println(&a)

	var b = "hello"
	fmt.Println(len(b))
	fmt.Println(&b)
}
