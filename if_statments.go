package main

import (
	"errors"
	"fmt"
)

func main(){

	// If
	input := 5
	if input % 2 == 0 {
		fmt.Println(input, "is even")
	}

	if input % 2 == 1 {
		fmt.Println(input, "is odd")
	}


	// if else
	input2 := 4
	if input2 % 2 == 0 {
		fmt.Println(input2, "is even")
	} else {
		fmt.Println(input2, "is odd")
	}


	input3 := 10
	// else if
	if input3 < 0 {
		fmt.Println("input can't be a negative number")
	} else if input3 % 2 == 0 {
		fmt.Println(input3, "is even")
	} else {
		fmt.Println(input3, "is odd")
	}

	// Initial if statements
	input4 := 21
	if err := validate(input4); err != nil {
		fmt.Println(err)
	} else if input4 % 2 == 0 {
		fmt.Println(input4,  "is even")
	} else {
		fmt.Println(input4, "is odd")
	}

}

func validate(input int) error {

	if input < 0 {
		return errors.New("input cant be negative")
	} else if input > 100 {
		return errors.New("input cant be over 100")
	} else if input % 7 == 0 {
		return errors.New("input cant be divisibe by 7")
	} else {
		return nil
	}
}
