package main

import (
	"fmt"
	"strings"
)


func main(){

	defer done()

	fizzBuzz(50)
	fmt.Println()
	itemsSold()

	fmt.Println()
	name, age := greeting()
	fmt.Printf("%s is %d", name, age)


	//Mapping index values to column headers
	fmt.Println()
	hdr := []string{"empid","employee",
		"address","hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"employee","empid",
		"hours worked","address","manager", "hourly rate"}
	csvHdrCol(hdr2)


	for i:=0;i<=15;i++{
		n,s:=fizzBuzz(i)
		fmt.Printf("Results: %v %v\n",n,s)
	}

	fmt.Println()
	variadicFunction("ee", 10,20,30)

	ff := []int{1,2,3}
	variadicFunction("a", ff ...)

	fmt.Println()
	i := []int{5,10,15}
	fmt.Println(sum(5,4))
	fmt.Println(sum(i...))

	// Ananymous functions
	func(){
		fmt.Println("Greeting")
	}()

	message := "Greeting"
	func(str string){
		fmt.Println(str)
	}(message)

	f := func(){
		fmt.Println("Executando dentro do anonimo")
	}

	fmt.Println("linha apos")
	f()

	j := 9
	x := func(i int)int {
		return i * i
	}

	fmt.Printf("the square of %d is %d\n",j,x(j))

	//Closures
	fmt.Println()
	fmt.Println()

	ii := 0
	incrementor := func() int {
		f := func() int {
			return 1
		}
		fmt.Println(f())
		ii += 1
		return ii
	}

	fmt.Println(incrementor())
	fmt.Println(incrementor())
	ii += 10
	fmt.Println(incrementor())

	fmt.Println(incr()())
	inc := incr()
	fmt.Println(inc())

	tt := 50
	fundec := decr(tt);

	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(fundec())
	fmt.Println(tt)

	// Function types
	calculator(add, 5, 6)
	calculator2(add2, 10, 3)
	calculator2(sub2, 10, 3)

	sq := square(9)
	fmt.Println(sq())
	fmt.Printf("Type of v: %T",sq)

	devS := salary(50, 2080, developerSalary)
	bossS := salary(150000, 25000, managerSalary)
	fmt.Println(devS, bossS)

	// defer




}

func done() {
	fmt.Println("Now I am done")
}

func salary(x,y int,f func(int,int)int)int{
	pay := f(x,y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(baseSalary, bonus int) int {
	return baseSalary * bonus
}

func square(x int) func() int {
	f := func() int {
		return x * x
	}
	return f
}

func calculator2(f func(int,int)int,i int , j int,){
	fmt.Println(f(i,j))
}

// Function types
type calc func(int,int) string

func add(i int, j int) string {
	r := i + j
	return fmt.Sprintf("Added %d + %d = %d", i,j,r)
}

func calculator(f calc, i, j int){
	fmt.Println(f(i,j))
}


func add2(i,j int) int {
	return i + j
}

func sub2(i,j int) int {
	return i - j
}



func decr(v int) func() int{

	return func() int {
		v--
		return v
	}

}




func incr() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func variadicFunction( s string,n ...int){
	fmt.Println(n)
}

func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i,v := range header{
		v = strings.TrimSpace(v)

		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}

	}
	fmt.Println(csvHeadersToColumnIndex)

}

func greeting() (name string, age int){
	name = "John"
	age = 21
	return
}

func itemsSold(){
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k,v := range items{
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v > 40 && v < 100{
			fmt.Println("meets expectations.")
		} else if v > 100 {
			fmt.Println("exceeded expectations.")
		}
	}

}

func fizzBuzz(i int) (int, string) {
	if i%15 == 0 {
		return i, "FizzBuzz"
	} else if i%3 == 0 {
		return i, "Fizz"
	} else if i%5 == 0 {
		return i, "Buzz"
	} else {
		return i, ""
	}
}

