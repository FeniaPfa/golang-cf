package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {

	john := User{Name: "John", Email: "john@doe.com", Active: true}

	jane := User{Name: "Jane", Email: "jane@doe.com", Active: true}

	johnStudent := Student{User: john, Id: "1f1"}

	fmt.Println(jane) // {Jane jane@doe.com true}
	fmt.Println(johnStudent) // {{John john@doe.com true} 1f1}

	fmt.Println(johnStudent.User.Name) // John

}