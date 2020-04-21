package main

import "fmt"

func main() {

	// padrao de pipeline uma coisa consumindo a aoutra

	numbers := generate(2,4,6,12)
	result := divide(numbers)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int  {

	channel := make(chan int)

	go func() {
		for _, n := range numbers{
			channel <- n
		}
	}()

	return channel
}

func divide(input chan int) chan int{

	channel := make(chan int)

	go func() {
		for n := range input {
			channel <- n /2
		}
		close(channel)
	}()

	return channel
}
