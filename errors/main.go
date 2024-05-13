package main

import (
	"errors"
	"fmt"
)

func divide(number , divider int)(int, error){
	if divider == 0 {
		return 0, errors.New("No es posible dividir sobre 0")
	} else{
		return number / divider, nil
	}
}

func main(){

	if result, err := divide(10, 0); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("El resultado es:", result)
	}


}