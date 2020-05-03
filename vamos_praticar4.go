package main

import "fmt"


const (
	logInfo = "Debug"
	idade = 14
)

func main(){

	fmt.Println("Hollo mundo")

	var fernando string = "nome"
	fmt.Println(fernando)

	fmt.Println(fmt.Sprintf("Meunome é %v e tenho %v anos", "fernando", 24))

	fmt.Println(logInfo, idade)

	arr := []string  {
		"nome",
		"ola",
	}

	fmt.Println(arr)
	fmt.Println(arr[1])

	for i,v := range arr{
		fmt.Println(i, v)
	}

	mapas := map[string]int {
		"idade":24,
		"altura":170,
	}

	fmt.Println(mapas["idade"])

	a := func(a,b int) int {
		return a + b
	}

	fmt.Println(calculo(a))
}

func calculo( a func(x,b int) int ) int {
	return a(50, 5)
}