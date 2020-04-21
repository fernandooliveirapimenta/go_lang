package main

import (
	"fmt"
	"sync"
)

func main() {

	channel := make(chan int)
	var waitG sync.WaitGroup
	waitG.Add(2)

	go func() {
		for i:=0;i<10;i++{
			channel <- i
		}
		waitG.Done()
	}()

	go func() {
		for i:=0;i<10;i++{
			channel <- i
		}
		waitG.Done()
	}()

	go func() {
		waitG.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}


}
