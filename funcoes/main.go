package main

import "fmt"

func f1(a,b int) int  {
	return a + b
}

func f2(a,b int) (int , string)  {
	a = a + b + 1
	b = a + 2
	return a + b , "ola"
}

func f3(a,b string) (x string)  {
	x = b
	return
}

func f4(numeros ...int)  {
	for _,v := range numeros {
		fmt.Println(v)
	}
}

func main() {

	fmt.Println(f1(5,10))

	a,b := f2(50,1)
	fmt.Println("a:",a, "b:",b)

	c := f3("50","nome")

	fmt.Println(c)


	f4(50,3,1,33)

	// Funcao anonima

	z := 0
	add := func() int {

		z += 2
		return z
	}
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
}
