package main

import (
	"fmt"
)

const (
	sunday = 0
	monday = 1
	tuesday = 2
	wednesday = 3
	thursday = 4
	friday = 5
	saturday = 6
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var level = "pkg"

func main() {

	fmt.Println(sunday, monday,
		tuesday, wednesday, thursday, friday, saturday)
	fmt.Println(Sunday, Monday,
		Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println("Main start: ", level)

	if true {
		fmt.Println("Block start: ", level)
		funcA()
	}

	level := 42
	fmt.Println("Shadow: ", level)
	funcA()

	// Message Bug

	count := 5
	message := ""

	if count > 5 {
		message = "Greater thar 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)


	// Bad count bug
	badCount := 0

	if badCount < 5 {
		badCount = 10
		badCount ++
	}

	fmt.Println(badCount == 11)

}

func funcA() {
	fmt.Println("funcA start: ", level)
}
