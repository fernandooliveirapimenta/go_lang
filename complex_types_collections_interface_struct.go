package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Definig array
	//arr := []int{1,2,3,4}
	//fmt.Println(arr)
	fmt.Printf("%v\n", defineArray())

	// Comparing arrays
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}:",comp1)
	fmt.Println("[5]int == [...]int{0,0,0,0,0}:",comp2)
	fmt.Println("[5]int == [9]int{0,0,0,0,9}:",comp3)


	// Initializing an Array using keys
	fmt.Println()
	i1,i2,i3 := initializingArraysKeys()
	fmt.Println("[10]int == [...]{9:0}:",i1)
	fmt.Println("[10]int == [10]{1,9:10,4:5}:",i2)
	fmt.Println("arr3:",i3)
	fmt.Println([]int{50:4})


	// Reading a single item from an array
	fmt.Println()
	fmt.Println(message())


	// Writing to an Array
	fmt.Println()
	fmt.Println(writingToArray())


	// Looping an Array
	fmt.Println()
	fmt.Println(loopingToArray())


	// Modifying contents of an array in a loop
	fmt.Println()
	var arra [10]int
	arra = modifyContentArray(arra)
	arra = modyfyContentArray2(arra)
	fmt.Println(arra)


	// Filing an array
	fmt.Println()
	var array [10]int
	for i := 1; i <= 10; i++ {
		array[i-1] = i
	}
	fmt.Println(array)


	// Working with slices
	fmt.Println()
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There wa an error")
	}

	// Activity Printing a user name based on user input
	fmt.Println()
	mapUsers := map[string]string{
		"305":"Sue",
		"204":"Bob",
		"631":"Jake",
		"073":"Tracy",
	}

	if len(os.Args) > 0 {
		arg := os.Args[1]
	    fmt.Println("Hi,", mapUsers[arg])
	}


	// Appendis multiple items to a slice
	locales := getLocals(getPassegArgs())
	fmt.Println("Locales to use:", locales)


	// Creting a locale checker
	if len(os.Args) > 0 {
		localeParts := strings.Split(os.Args[1], "_")
		if len(localeParts) == 2 {
			parsed := locale{
				territory: localeParts[1],
				lenguage: localeParts[0],
			}

			if !localeExists(parsed) {
				fmt.Printf("Locale not supported: %v\n", os.Args[1])
			} else {
				fmt.Println("Locale passed is supported")
			}
		}
	}

	// Creating slices from a slice
	fmt.Println()
	fmt.Println(creatingSliceFromSlice())


	// Using make to control the capacity of a slice
	fmt.Println()
	s1,s2,s3 := getSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))

	// Controlling internal slice behavior
	fmt.Println()
	l1,l2,l3 := linked()
	fmt.Println("Linked:",l1,l2,l3)
	nl1, nl2 := notLink()
	fmt.Println("No Link:", nl1, nl2)
	cl1,cl2 := capLinked()
	fmt.Println("Cap Link:", cl1, cl2)
	cnl1,cnl2 := capNoLink()
	fmt.Println("Cap No Link:", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)




}

func appendNoLink() (int,int) {
	s1 := []int{1,2,3,4,5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink()(int,int,int){
	s1 := []int{1,2,3,4,5}
	s2 := make([]int, len(s1))
	copied := copy(s2,s1)
	s1[3] = 99
	return s1[3],s2[3],copied
}

func capNoLink()(int,int){
	s1 := make([]int,5,10)
	s1[0],s1[1],s1[2],s1[3],s1[4] = 1,2,3,4,5
	s2 := s1
	s1 = append(s1, []int{10:11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked()(int,int){
	s1 := make([]int,5,100)
	s1[0],s1[1],s1[2],s1[3],s1[4] = 1,2,3,4,5
	s2 := s1
	s1 = append(s1,6)
	s1[3] = 99
	return s1[3],s2[3]
}

func notLink() (int,int){
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3],s2[3]
}

func linked() (int,int,int) {
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99
	return s1[3],s2[3],s3[3]
}

func getSlices()([]int,[]int,[]int){
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1,s2,s3
}

func creatingSliceFromSlice() string{
	s := []int{1,2,3,4,5,6,7,8,9}
	m := fmt.Sprintln("First: ", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last: ", s[len(s)-1], s[len(s)-1:], s[len(s)-1])
	m += fmt.Sprintln("First 5:",s[:5])
	m += fmt.Sprintln("Last 4:",s[5:])
	m += fmt.Sprintln("Middle 5:",s[2:7])
	return m
}

type locale struct {
	lenguage string
	territory string
}

func getLocalesStruct() map[locale]struct{}{
	suported := make(map[locale]struct{},5)

	suported[locale{"en", "US"}] = struct{}{}
	suported[locale{"en", "CN"}] = struct{}{}
	suported[locale{"fr", "CN"}] = struct{}{}
	suported[locale{"fr", "FR"}] = struct{}{}
	suported[locale{"ru", "RU"}] = struct{}{}

	return suported
}

func localeExists(l locale) bool  {
	_, exists := getLocalesStruct()[l]
	return exists
}


func getPassegArgs() []string{
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func getLocals(extraLocals []string) []string{
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}

func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func getPassedArgs(mingArgs int) []string{
	if len(os.Args) < mingArgs {
		fmt.Printf("At least %v arguments are needed\n", mingArgs)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func modifyContentArray(arr [10]int) [10]int{
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func modyfyContentArray2(arr [10]int) [10]int{

	for i := 0; i < len(arr); i ++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func loopingToArray() string {
	m := ""
	arr := [4]int{1,2,3,4}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}

func writingToArray() string  {
	arr := [4]string{"ready", "Get", "Go", "to"}
	arr[1] = "Its"
	arr[0] = "time"
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])

}

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}

	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
func initializingArraysKeys() (bool, bool, [10]int){
	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1,9:10,4:5}
	return arr1 == arr2, arr1 == arr3, arr3
}

func compArrays() (bool, bool, bool) {
  var arr1 [5]int
  arr2 := [5]int{0}
  arr3 := [...]int{0,0,0,0,0}
  arr4 := [5]int{0, 0, 0, 0, 9}

  return arr1 == arr2, arr1 == arr3, arr1 == arr4

}

func defineArray() [10]int  {
	var arr [10]int
	return arr
}