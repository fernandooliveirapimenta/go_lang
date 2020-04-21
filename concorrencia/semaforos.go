package main

import "fmt"

func main() {
	//semaforos podem substituir os waiting groups
	channel := make(chan int)
	ok := make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			channel <- i
		}
		ok <- true
	}()

	go func() {
		for i:= 0; i <10;i++{
			channel <-i
		}
		ok <- true
	}()

	go func() {
		<-ok
		<-ok
		close(channel)
	}()

	for n := range channel{
		fmt.Println(n)
	}

}
