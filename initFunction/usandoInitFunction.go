package main

import "fmt"

var name = "Gopher"

var budgetCategories = make(map[int]string)



func init(){
	fmt.Println("Hello, ", name)
	fmt.Println("initializing out budgetCategories")

	budgetCategories[1] = "Car Insurance"
	budgetCategories[2] = "Mortgade"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[7] = "Groceries"
	budgetCategories[8] = "Car Payment"


}

func init(){
	fmt.Println("Second")
}

func init(){
	fmt.Println("Third")
}

func main() {

	fmt.Println("Main")


	for key, value := range budgetCategories {
		fmt.Printf("Key: %v, value: %v \n", key, value)
	}
}
