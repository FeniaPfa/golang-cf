package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}
// to be type interface must implement all the methods with the same names and params
type Dog struct {
	Name string
}

func (self Dog) Eat(){
	fmt.Println("Dog is eating")
}
func (self Dog) Sleep(){
	fmt.Println("Dog is sleeping")
}

	func execAction(animal Animal){
		animal.Eat()
		animal.Sleep()
	}

func main() {

	myDog := Dog{Name: "Dogg"}

	fmt.Println(myDog)

	execAction(myDog)

}