package main

import "fmt"

// es global
var username string // ""

func function1(){
	username = "Func 1"
	fmt.Println("Func 1:", username)
}
func function2(){
	username = "Func 2"
	fmt.Println("Func 2:", username)
}

func main(){

	
	username = "Cody"
	fmt.Println(username)
	
	function1()
	function2()

	fmt.Println(username)
}