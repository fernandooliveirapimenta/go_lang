package main

import (
	"fmt"
	"time"
)

func main() {

	dayBorn := time.Monday

	switch dayBorn {

	case time.Monday:
		fmt.Println("Mondays child is fair of face")
	case time.Tuesday:
		fmt.Println("Tuesday child is full of grace")
	case time.Wednesday:
		fmt.Println("Wednesday child is full woe")
	case time.Thursday:
		fmt.Println("Thusday child has far to go")
	case time.Friday:
		fmt.Println("Friday is child is loving and giving")
	case time.Saturday:
		fmt.Println("Saturday child works hard for a living")
	case time.Sunday:
		fmt.Println("Sunday child is bonny and blithe")
	default:
		fmt.Println("Error day born not valid")
	}


	// Multiple valures

	switch dayBorn {
	case time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday:
		fmt.Println("Born os a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Error, day born not valid")

	}


	// Expressinless switch statements

	switch dayBorn2 := time.Sunday;{
	case dayBorn2 == 50:
		fmt.Println("50")
	case dayBorn2 == time.Sunday || dayBorn2 == time.Saturday:
		fmt.Println("Born on the weekend ")
	default:
		fmt.Println("Born some other day")

	}
}
