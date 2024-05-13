package main

import "fmt"

type Operation func(num1, num2 int) int

func sum(number1 , number2 int) int {
	return number1 + number2
}

func sub(number1 , number2 int) int {
	return number1 - number2
}

func doOperation(function Operation, num1 , num2 int){
	fmt.Println("Before the operation")

	result := function(num1, num2)
	fmt.Println("Result:", result)

	fmt.Println("After the operation")
}

// named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){

	doOperation(sum, 20, 30)
	doOperation(sub, 30, 20)

	fmt.Println(split(17))

}
