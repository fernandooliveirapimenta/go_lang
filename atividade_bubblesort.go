package main

import "fmt"

func main() {
	entrada := []int{5,8,2,4,0,1,3,7,9,6}
	fmt.Println(len(entrada))
	fmt.Println("Before:", entrada)

	for i := 0; i < len(entrada) - 1 ; i++  {
		for j := 1; j < len(entrada) - i; j++ {
			if entrada[j-1] > entrada[j] {
				tmp := entrada[j-1]
				entrada[j-1] = entrada[j]
				entrada[j] = tmp
			}
		}
	}

	fmt.Println(entrada)
}
