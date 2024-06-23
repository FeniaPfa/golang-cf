package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createUser() {
	clearConsole()
	fmt.Println("Enter username:")
	username := readLine()
	fmt.Println("Enter email:")
	email := readLine()
	fmt.Println("Enter age:")
	age, err := strconv.Atoi(readLine())
	if err != nil {
		panic("Unable to convert string to int")
	}

	id++
	user := User{id, username, email, age}

	users[id] = user
	fmt.Println(">>>> User created \n")
	fmt.Println(users)
}

func getUsers() {
	clearConsole()

	fmt.Println("User list")
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
	fmt.Println("\n")
}

func updateUser() {
	fmt.Println("Updated user")
}

func deleteUser() {
	clearConsole()
	fmt.Println("Enter user ID")
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("unable to convert id to int")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println(">>> User deleted \n")
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("Itsnot posible get the value")
	} else {

		return strings.TrimSuffix(strings.TrimSpace(option), "\n")
	}

}

func main() {
	users = make(map[int]User)

	for {
		var option string

		reader = bufio.NewReader(os.Stdin)

		fmt.Println("A) Create")
		fmt.Println("B) GET")
		fmt.Println("C) Update")
		fmt.Println("D) Delete")

		fmt.Println("Enter a valid option")

		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "create":
			createUser()
		case "b", "get":
			getUsers()
		case "c", "update":
			updateUser()
		case "d", "delete":
			deleteUser()
		default:
			fmt.Println("No valid option", option)
		}
	}

	fmt.Println("Quitting...")
}
