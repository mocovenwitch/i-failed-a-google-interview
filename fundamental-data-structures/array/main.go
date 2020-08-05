package main

import "fmt"

func main() {
	tryArray()
}

func tryArray() {
	a := [5]int{1, 2, 3}
	fmt.Println(a)

	// $go run main.go
	// [1 2 3 0 0]
}
