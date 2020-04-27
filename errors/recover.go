package main

import (
	"errors"
	"fmt"
)

func main() {

	a()
	fmt.Println("This line will get printed from main() funcion")

}

func a(){
	b("good-bye")
	fmt.Println("Back in function a()")
}

func b(s string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error in func b()",r)
		}
	}()

	if s == "good-bye" {
		panic(errors.New("something went wrong"))
	}

	fmt.Println(s)
}
