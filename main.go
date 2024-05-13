package main

import "fmt"

// go build main.go -> Archivo ejecutable
// go run main.go

// const se declara fuera de la funcion principal
const Curso string = "Curso"

func main(){
	// Curso = "Nuevo valor" // <- no se puede porque es una constante

	var nombre string

	var edad int

	nombre = "Nombre"
	edad = 1

	fmt.Println(nombre)
	fmt.Println(edad)

	var nombre2 string = "Nombre 2"
	fmt.Println(nombre2)

	// sin indicar el tipo de dato, igual no puede cambiar de tipo
	// compilador intuye el tipo de dato
	variable := "string"
	fmt.Println(variable)

	var altura = 2.2 // float
	fmt.Println(altura)

	// declarar 3 variables al mismo t iempo
	// var nombre3, apellido, pais string

	// var nombre3, apellido, pais = "nombre" , "apellido" , "pais"

	// agruparlos por tipo
	nombre3, apellido, pais := "nombre" , "apellido" , "pais" 
	edad2, altura2 := 8 , 12

fmt.Println(nombre3, apellido, pais, edad2, altura2)






}

