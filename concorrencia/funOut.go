package main

import "fmt"

func main() {

	c := genereteFunOut(5,10)

	d1 := divideFunOut(c)
	d2 := divideFunOut(c)

	fmt.Println(<- d1)
	fmt.Println(<- d2)


}

func genereteFunOut(nums ...int) chan int {
	channel := make(chan int)
	go func() {
		for _,n:=range nums{
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func divideFunOut(input chan int) chan int{
	channel := make(chan int)
	go func() {
		for nums := range input{
			channel <- nums/2
		}
		close(channel)
	}()
	return channel
}