package main

import "fmt"

func main() {
	// type, length, cap
	slice := make([]int, 2, 3)

	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 3
	fmt.Println(slice) // [0 ,0]
}