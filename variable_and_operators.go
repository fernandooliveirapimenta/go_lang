package main

import (
	"fmt"
	"time"
)

var defaultOffset = 10

func main(){
	debug := false
	logLevel := "info"
	startUpTime := time.Now()

	fmt.Println(debug, logLevel, startUpTime)

	//Delcaring multiple variables with a short variable
	calvo, idade, anoAtual := false, 24, time.Now().Year()
	fmt.Println(calvo, idade, anoAtual)


	//Declaring multiple variables from a function
	x, y, v := getConfig()
	fmt.Println(x, y, v)


	//Using var to declare multiple variables in one line

		// Type only
		var start, middle, end float32
		fmt.Println(start, middle, end)

		// Init value mixed type
		var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
		fmt.Println(name, left, right, top, bottom)

		// works with functions also
		var a, b, c = getConfig()
		fmt.Println(a, b, c)

	// Non-english variable names
	デバッグ := false
	日志级别 := "info"
	ይጀምሩ := time.Now()
	_A1_Μείγμα := ""
	fmt.Println(デバッグ, 日志级别, ይጀምሩ, _A1_Μείγμα)

	// Changing the value of a variable
	offset := 5
	fmt.Println(offset)

	offset = 10
	fmt.Println(offset)

	offset2 := defaultOffset
	fmt.Println(offset2)

	offset2 = offset2 + defaultOffset
	fmt.Println(offset2)

	// Changing multiple values at once
	query, limit, offset := "bat", 10, 0
	query, limit, offset = "ball", offset, 20
	fmt.Println(query, limit, offset)


	// Using operators with numbers
	var total float64 = 2 * 13
	fmt.Println("Sub:", total)

	total = total + (4 * 2.25)
	fmt.Println("Sub: ", total)

	total = total - 5
	fmt.Println("Sub: ", total)

	tip := total * 0.1
	fmt.Println("Tip: ", tip)

	split := total/2
	fmt.Println("Split: ", split)

	visitCount := 24
	visitCount = visitCount + 1

	remainger := visitCount % 5
	if remainger == 0 {
		fmt.Println("With this visit, you earned a reward.")
	}

	givenName := "John"
	familyName := "Smith"
	fullName := givenName + " " + familyName
	fmt.Println("Hello, ", fullName)


	// Shorthand operator
	count := 5
	count += 5
	fmt.Println(count)

	count ++
	fmt.Println(count)

	count --
	fmt.Println(count)

	count -= 5
	fmt.Println(count)


	name2 := "John"
	name2 += " Smith"
	fmt.Println("Hello, ", name2)

	// Comparing values
	alvo := 15

	fmt.Println("First visit: ", alvo == 1)
	fmt.Println("Return visit: ", alvo != 1)
	fmt.Println("Silver member: ", alvo >= 10 && alvo < 21)
	fmt.Println("Gold member: ", alvo > 20 && alvo <= 30)
	fmt.Println("Platinum member: ", alvo > 30)

	// Zero values
	var count2 int
	fmt.Printf("Count: %v \n", count2)
	var discount2 float64
	fmt.Printf("Discount: %v \n",discount2)
	var debug2 bool
	var msg string
	var emails []string
	var startTime time.Time
	fmt.Println(debug2, msg, emails, startTime)

}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
