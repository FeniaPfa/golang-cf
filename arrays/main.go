package main

import "fmt"

func main() {
	// arrays length is static
	var numbers [5]int // <-- declare
	fmt.Println(numbers) // [0 0 0 0 0] <-- start value as int

	//var[index]
	numbers[0] = 100 // can access by index
	fmt.Println(numbers) // [100 0 0 0 0] 

	// fmt.Println(numbers[6]) // <-- error
	// numbers[1] = "string value" // <-- also error

// asign values when declaring
					// [length] type {...values}
numbers2 := [5] int {100, 200, 300, 400, 500}
fmt.Println(numbers2) // [100 200 300 400 500]

// Can asign especific index order

					// compiler infers length
currencies := [...] string {0: "CLP", 2: "Dolar", 1: "Euro"}
fmt.Println(currencies) // [CLP Euro Dolar]

subCurrencies := currencies[0: 2] // this is a slice
fmt.Println(subCurrencies) // [CLP Euro]

}