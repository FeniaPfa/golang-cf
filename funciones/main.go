package main

import "fmt"

func greeting(name string){
	fmt.Println("Hello " , name)
}

// like ts
// or when are the same type (number1, number2 int) int
func sum (number1 int, number2 int) int {
	return number1 + number2
}

func sum2 (num1 , num2 int) (int, string){
	return num1 + num2, "sum"
}

func main(){

	greeting("Test")

	result := sum(1 , 2)
	result2, myString := sum2(1,4)
	fmt.Println(result)
	fmt.Println(result2, myString)


	// IIFE
	func(){
		fmt.Println("unnamed function")
	}()

	myFunc := func(){
		fmt.Println("Anonymous Function")
	}
	
	myFunc() // Anonymous Function
}