package main

import (
	"fmt"
	"time"
)

func main() {
	// mandar informacoes para diferentes processos
	// canal de comunicacao entre os processos

	// existe chanel com buffe e sem bufer
	channel := make(chan int)

	go func() {
		for i:=0;i<10;i++{
			channel <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<- channel)
		}
	}()

	time.Sleep(time.Second)


}
