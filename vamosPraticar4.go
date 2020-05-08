package main

import (
	"fmt"
	"strings"
)

type Garrafa struct {
	Nome string
	Qtd int
}

type Caixa struct {
	Litros int
	Garrafas []Garrafa
}

func (c *Caixa) ContarTipo(tipo string) (int, string) {

	f := func (v, a string) bool {
		return strings.EqualFold(v, a);
	}

	var aux int
	for _, v := range c.Garrafas {
		if f(v.Nome, "coca" ) {
			aux ++
		}
	}

	c.Litros = c.Litros + aux

	return aux, "sucesso"
}

func main() {


	caixa := Caixa{
		Litros: 2,
		Garrafas: []Garrafa{
			{
				Nome: "Coca",
				Qtd:1,
			},
			{
				Nome: "Coca",
				Qtd:1,
			},
			{
				Nome: "Coca",
				Qtd:1,
			},
			{
				Nome: "Guarana",
				Qtd:1,
			},
		},
	}
	qtd, stado := caixa.ContarTipo("coca")
    caixa.ContarTipo("coca")
	fmt.Printf("qtd coca %v , %v \n", qtd, stado )

	fmt.Println(caixa)
	for i := 0; i < 10; i++ {
		//for ii, v := range caixa.Garrafas {
		//	fmt.Printf("Indice: %v valor: %v", ii,v)
		//}
		if i % 2 == 0 {
			fmt.Printf(" É par : %v \n",i)
		} else {
			fmt.Println("impar")
		}
	}

}
