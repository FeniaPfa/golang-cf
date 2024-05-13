package main

import "fmt"

// const (
// 	Domingo int = 0
// 	Lunes int = 1
// 	Martes int = 2
// 	Miercoles int = 3
// 	Jueves int = 4
// 	Viernes int = 5
// 	Sabado int = 6
// )

const (
		Domingo int = iota + 1 // indica que los siguientes valores son una secuencia desde 0
		Lunes
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
	)

func main(){
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}