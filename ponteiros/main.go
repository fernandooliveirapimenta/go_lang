package main

import "fmt"

// essa funcao esta mudando o valor direto na memoria ram
func xpto(a *int) int {
	*a = *a + 1
	return *a
}

func main() {

	b := 10

	fmt.Println("xpto:",xpto(&b))
	fmt.Println("b:",b)

	x := 10
	fmt.Println("x:",x)
	fmt.Println("x:",&x) // pegando referencia dele na memori usa &

	y := &x // y agora esta com a referencia de x
	fmt.Println("y:",y)

	fmt.Println("y:",*y) // o * busca na memoria o valor desse ex: 0xc820..

	*y = 20 // estou mudando o valor para onde o y estava apontando que é o mesmo endereco do x

	fmt.Println("x:",x)

	var z *int = &x
	fmt.Println("z:", z)
}
