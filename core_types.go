package main

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"unicode"
)

func main(){
	// Boolean
	fmt.Println(10 > 5)
	fmt.Println(10 == 5)


	// Medir complexidade da senha
	fmt.Println()
	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	if passwordChecker("This!I5A") {

		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	// Numbers
	fmt.Println()
	var list []int8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Total Alloc (Heap) = %v MiB \n", m.TotalAlloc/1024/1024)

	// Floating-Point Number Accuracy
	fmt.Println()
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a/3)
	fmt.Println(b/3)
	fmt.Println(c/3)
	fmt.Println(a/3 * 3)
	fmt.Println(b/3 * 3)
	fmt.Println(c/3 * 3)

	// Overflow and wraparound
	fmt.Println()
	//var aa int8 = 128 // max 128
	//fmt.Println(aa)

	var aaa int8 = 125
	var bbb uint8 = 253

	for i := 0; i < 5; i++ {
		aaa++
		bbb++
		fmt.Println(i, ")", "int8", aaa, "uint8", bbb)
	}

	// Big numbers
	fmt.Println()
	intA := math.MaxInt64
	intA = intA + 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64:",math.MaxInt64)
	fmt.Println("Int:",intA)
	fmt.Println("Big Int:",bigA.String())


	// Byte
	fmt.Println()

	// Text
	fmt.Println()
	comment1 := `This is the BEST thing ever!`
	comment2 := `This is the BEST \nthing ever!`
	comment3 := "This is the BEST \nthing ever!"
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	fmt.Println()
	cmd := `In "Windows" the user directory is "C:\Users\"`
	cmd2 := "In \"Windows\" the user directory is \"C:\\Users\\\""
	fmt.Println(cmd)
	fmt.Println(cmd2)


	// Rune
	fmt.Println()
	username := "Sir_King_Über"
	// somente bytes
	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}

	// converter novamente para strings
	fmt.Println()
	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}

	// utilizando rune
	fmt.Println()
	runes := [] rune(username)
	for i := 0; i < len(runes); i++ {
		//fmt.Print(runes[i], " ")
		fmt.Print(string(runes[i]), " ")
	}

	// Safely Looping over a string
	fmt.Println()
	logLevel := "デバッグ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}

	// Length of a string
	fmt.Println()
	fmt.Println("Bytes:",len(username))
	fmt.Println("Runes:",len([]rune(username)))
	// Limit to 10 charaters
	fmt.Println(string(username[:10]))
	fmt.Println([]rune(username)[:10])
	fmt.Println(string([]rune(username)[:10]))


	// The nil value
	fmt.Println()
	var message *string
	if message == nil {
		fmt.Println("error, unexpected nil value")
	}
	fmt.Println(&message)



}


func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumver := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumver = true
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumver && hasSymbol
}