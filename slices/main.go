package main

import "fmt"

// Slices are dynamic not like arrays

func main(){
						// nothing is in []
	numbers := [] int { 1, 2 ,3 ,4 ,5}

		// append returns a new slices it doesnt modify the original so must be stored in a variable
	//	append(numbers , 6) <-- doesnt work
	numbers = append(numbers, 6) // same as .push(x)
	numbers = append(numbers, 7)
	numbers = append(numbers, 8) 
	numbers = append(numbers, 9) 
	numbers = append(numbers, 10) 

	
	fmt.Println(numbers)
	
	
	newSlice := numbers[0:5]
	fmt.Println(newSlice) // [1 2 3 4 5]


	numbers[0] = 100
	fmt.Println(numbers) // [100 2 3 4 5 6 7 8 9 10]
	fmt.Println(newSlice) // [100 2 3 4 5]

	meses := [] string {"Enero", "Febrero", "Marzo" , "Abril", "Mayo" , "Junio", "Julio" , "Agosto", "Septiembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("Long : %v, Cap: %v %p\n", longitud, capacidad, meses) // Long : 9, Cap: 9  0xc00007a000


	meses = append(meses, "Octubre") // Si cap esta full se genera un nuevo arreglo
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")


	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("Long : %v, Cap: %v %p\n", longitud, capacidad, meses) // Long : 12, Cap: 18 0xc000088000 <-- apunta a otro espacio de memoria

}