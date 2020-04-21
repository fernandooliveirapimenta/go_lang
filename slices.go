package main

import "fmt"

func main() {

	//slice nao tem tamanho fixo somente tamanho inicial
	// deixa o array mais dinamico
	slice := make([]int, 2,5)
	//slice = append(slice, 10, 2, 5, 40)

	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))

	//for i:=0;i<20;i++{
	//	slice = append(slice, i)
	//	fmt.Println("Len: ", len(slice),  "Cap: ", cap(slice))
	//}


	fmt.Println()

	sTeste := slice

	slice = append(slice, 1, 2, 3, 5, 7)
	slice[0] = 10

	fmt.Println(slice)

	fmt.Println(sTeste)

	fmt.Println()
	ss := []string {
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(ss)
	fmt.Println(ss[:2])
	fmt.Println(ss[1:2])
	fmt.Println(ss[2:4])
	fmt.Println(ss[2:])


}
