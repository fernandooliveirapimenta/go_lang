package main

import "fmt"

func main() {

	for i := 1 ; i <= 100; i++ {

		divisivel3 := i % 3 == 0
		divisivel5 := i % 5 == 0
		if divisivel3 && divisivel5 {
			fmt.Println("FizzBuzz")
		} else if divisivel3 {
			fmt.Println("Fizz")
		} else if divisivel5 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
