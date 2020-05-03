package main

import "fmt"

var bc = make(map[int]string)
var ptc = make(map[string]int)

func init(){
	fmt.Println("Initializing our bc")
	bc[1] = "Car Insurance"
	bc[2] = "Mortgage"
	bc[3] = "Electricity"
	bc[4] = "Retirement"
	bc[5] = "Vacation"
	bc[7] = "Groceries"
	bc[8] = "Car Payment"
}

func init(){

	fmt.Println("Assign our Payees to categories")
	ptc["Nationwide"] = 1
	ptc["BBT Loan"] = 2
	ptc["First Energy Electric"] = 3
	ptc["Ameriprise Financial"] = 4
	ptc["Walt Disney World"] = 5
	ptc["ALDI"] = 7
	ptc["Martins"] = 7
	ptc["Wal Mart"] = 7
	ptc["Chevy Loan"] = 8

}

func main() {
	fmt.Println("In main, printing p t c")

	for k,v := range ptc {
		fmt.Printf("Payee: %s, Category: %s \n", k, bc[v])
	}
}
