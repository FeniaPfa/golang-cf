package main

import "fmt"

type User struct {
	Name  string
	Email string
}
			// also could be this or some name that makes sense
func (self *User) setName(name string){
	self.Name = name
}

func (self *User) getName() string{
	return self.Name
}

func main() {

	cody := User{Name: "Cody", Email: "cody@mail.com"}
	fmt.Println(cody) // {Cody cody@mail.com}
	
	cody.setName("Juan")
	fmt.Println(cody) // {Juan cody@mail.com}

	fmt.Println(cody.getName()) // Juan

}