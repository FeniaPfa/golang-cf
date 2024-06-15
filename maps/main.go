package main

import "fmt"

func main() {
	days := make(map[int]string)

	days[0] = "Domingo"
	days[1] = "Lunes"
	days[2] = "Martes"
	days[3] = "Miercoles"
	days[4] = "Jueves"
	days[5] = "Viernes"
	days[6] = "Sabado"

	days[4] = "JUEVES"

	fmt.Println(days) // map[0:Domingo 1:Lunes 2:Martes 3:Miercoles 4:JUEVES 5:Viernes 6:Sabado]
	delete(days, 1)
	
	fmt.Println(days) // map[0:Domingo 2:Martes 3:Miercoles 4:JUEVES 5:Viernes 6:Sabado]

	grades := make(map[string] []int)

	grades["Nombre"] = []int {1, 2 ,3 ,4}
	fmt.Println(grades) // map[Nombre:[1 2 3 4]]

	users := map[int] string{} // make

	users[0] = "User 1"
	users[1] = "User 2"
	users[2] = "User 3"
	users[3] = "User 4"

	fmt.Println(users)
	
	for key, value := range users{
		fmt.Println(key, value)
	}
	// 0 User 1
	// 1 User 2
	// 2 User 3
	// 3 User 4




}