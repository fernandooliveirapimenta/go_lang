package main

import (
	"fmt"
	"time"
)

func main() {
	// var <name> *<type>
	// ou  <name> := new(<type>)

	// <var1> := &<var2>

	// Getting a pointer
	var count1 *int
	count2 := new(int)
	countTemp := 5

	count3 := &countTemp

	t := &time.Time{}

	fmt.Printf("count1: %v \n", count1)
	fmt.Printf("count2: %v \n", count2)
	fmt.Printf("count3: %v \n", count3)
	fmt.Printf("time: %v \n", t)

	fmt.Println(*count3, countTemp)

	*count3 = 10

	fmt.Println(*count3, countTemp)

	// Getting a value from a pointer
	fmt.Println()
	fmt.Println()
	var tt *int
	passaro := new(int)
	countPassaro := 5
	tribo := &countPassaro
	tempo := &time.Time{}

	if tt != nil {
		fmt.Printf("tt: %v \n", *tt)
	}

	if passaro != nil {
		fmt.Printf("passari: %v \n", *passaro)
	}

	if tribo != nil {
		fmt.Printf("tribo: %v \n", tribo)
		fmt.Printf("tribo: %v \n", *tribo)
	}

	if tempo != nil {
		fmt.Printf("tempo: %v \n", *tempo)
		fmt.Printf("tempo: %v \n", tempo)
		fmt.Printf("tempo: %v \n", tempo.String())
	}

	// Function design with pointers
	var count int
	add5Value(count)
	fmt.Println("add5Value post: ", count)

	add5ValuePoint(&count)
	fmt.Println("add5Point post: ", count)


	// Pointer value swap

	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)



}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5ValuePoint(count *int) {
	*count += 5
	fmt.Println("add5Point: ", *count)
}

func swap(a *int, b *int){
	aux := *a
	//
	*a = *b

	*b = aux
}