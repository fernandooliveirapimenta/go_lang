package main

import "fmt"

func main() {
	// 1
	channel := make(chan  int)

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}


	// 2
	fmt.Println()
	x := make([]int, 2)
	x[0], x[1] = 1, 2
	for i,v := range x {
		fmt.Println(bool(&x[i] == &v))
	}


	// 7
	fmt.Println()
	os := map[string]int{
		"mac" : 3,
		"windows" : 1,
		"unix" : 0,
		"linux" : 2,
	}
	for k, v := range os{
		fmt.Println("os", v, "Key", k)
	}

	// 8
	fmt.Println()
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}


	// 12
	fmt.Println()
	if count := 10; count > 5 {
		fmt.Println("Count is greater than 5")
	}


	// 14
	fmt.Println()
	ii := 2
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("dafault")

	}

	iii := 2
	switch {

	case iii == 1:
		fmt.Println("one")
	case iii == 2:
		fmt.Println("two")
	default:
		fmt.Println("default")

	}

	// 17
	//fmt.Println()
	//for {break}
	//for iiii:= 0; { fmt.Println(iiii) break }
	//
	//for true { break}
}
