package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buff"
	messages <- "ch"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
