package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Recieved message ", msg)
	default:
		fmt.Println("No messages recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no messages recieved")
	}

	select {
	case msg := <-messages:
		fmt.Println("Recieved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signals", sig)
	default:
		fmt.Println("No activity")
	}

}
