package main

import (
	"errors"
	"fmt"
)

func main() {
	//nums := []int{1,2,3}

	//for i:=0;i<10;i++{
	//	fmt.Println(nums[i])
	//}

	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()
	msg := "good-bye"
	message(msg)
	fmt.Println("This line will not get printed")
}

func message(msg string){
	if msg == "good-bye" {
		panic(errors.New("Something went wrong"))
	}
}
