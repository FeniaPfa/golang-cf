package main

import (
	"fmt"
	// "reflect"
)

func main(){
	// var curso string = "Curso"
	// var curso = "Curso"
	curso := "Curso"

	fmt.Println(curso) 

	cursoLength := len(curso) // int

	fmt.Println(cursoLength)

	firstChar := curso[0] // Char -> uint8
	lastChar := curso[ len(curso) - 1]



	fmt.Println(firstChar)
	// fmt.Println(reflect.TypeOf(firstChar)) // uint8

	// En windows el antivirus no los deja
	fmt.Printf("%c^\n", firstChar) // C
	fmt.Printf("%c^\n", lastChar) // o


	v := "42" // 
	fmt.Printf("v is of type %T\n", v) // v is of type string


}