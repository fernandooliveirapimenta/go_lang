package main

import "fmt"

func main() {

	// é obrigratorio channel rodar em uma goroutine
	channel := make(chan int)
	//channel <- 10 deadlock
	go func() {
		channel <- 10
	}()
	fmt.Println(<- channel)


}
