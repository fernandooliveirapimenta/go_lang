package main

import "fmt"

func main() {

	var x [10]int
	fmt.Println(x)

	x[0] = 10
	x[3] = 10
	fmt.Println(len(x))
	fmt.Println(x)

	fmt.Println(x[4])


}
