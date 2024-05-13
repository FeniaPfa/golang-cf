package main

import "fmt"

type Operation func(balance, quantity int) int

func createOperation(typo string) Operation {
	if typo == "sum" {
		return func(bal, qty int) int {
			return bal + qty
		}
	} else if typo == "sub" {
		return func(bal, qty int) int {
			return bal - qty
		}
	} else {
		return func(bal, qty int) int {
			return bal * qty
		}
	}
}

func main() {

	sum := createOperation("sum")

	result := sum(40, 50)

	fmt.Println(result)

}