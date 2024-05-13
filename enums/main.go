package main

import "fmt"

type UserType int

// create constants from a type

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {

	john := User{Username: "John", Type: Student}
	jane := User{Username: "John", Type: Teacher}

	fmt.Println(john)
	fmt.Println(jane)

	if john.Type == Student {
		fmt.Println("User" , john.Username, "is a student")
	}

}