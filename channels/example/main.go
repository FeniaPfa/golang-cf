package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"

	fmt.Println(len(c), cap(c)) // length capacity 2 2

	close(c) // closes the channel

	for channelMessage := range c {
		fmt.Println(channelMessage)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Message 1", email1)
	go message("Message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
