package main

import (
	"fmt"
	"unsafe"
)

func main() {
	tryArray()
	tryArrayLocation()
}

func tryArray() {
	a := [5]int{1, 2, 3}
	fmt.Println(a)

	// $go run main.go
	// [1 2 3 0 0]
}

func tryArrayLocation() {
	a := [3]int{1, 2, 3}

	for i, element := range a {
		fmt.Println(element, " at memory:", &a[i])
	}

	// 1  at memory: 0xc00011a000
	// 2  at memory: 0xc00011a008
	// 3  at memory: 0xc00011a010

	// Off topic, we can see an int element occupies 8 bytes, then prove it
	fmt.Println(unsafe.Sizeof(a[0]))
}
