package main

import "fmt"

func main() {
	// mandar informacoes para diferentes processos
	// canal de comunicacao entre os processos

	// existe chanel com buffe e sem bufer
	msg := make(chan string)

	go func() {
		msg <- "Hello world"
	}()

	result := <- msg
	fmt.Println(result)

}
