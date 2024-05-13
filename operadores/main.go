package main

import "fmt"

func main(){
	fmt.Println("Operadores")

	// ==
	// !=
	// >
	// <
	// >=
	// <=

	// var numero = 10
	// var resultado = numero > 5 // true
	var resultado = "Cody" == "Cody" // true

	fmt.Println(resultado)

	fmt.Println("Operadores lógicos")
	// && || !
	// result := 20 == 20 && 30 == 30 // true

	// result := 20 == 20 || 30 == 1 // true

	// result := 15 == 15 && 60 < 100 && (90 < 100 || 100 == 90) // true

	result := !true // false

	fmt.Println(result)
}