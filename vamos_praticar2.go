package main

import "fmt"

type Pessoa struct {
	Idade int
	Nome string
}

type Fernando struct {
	Pessoa
	EhBonito bool
}

func (f *Fernando) falar()  {
	fmt.Println("Fernando falando")
}

type TT interface {
	falar()
}

func main(){

	var nome string = "Fernando"
	apelido := "Pimenta"

	fmt.Println(nome)
	fmt.Printf("Nome : %s, apelido %v \n", nome, apelido)

	tipos := []string {"neca", "app", "boy", "she"}

	fmt.Println(tipos)

	tipos = append(tipos, "fernando")

	fmt.Println(tipos[2:])


	mp := map[string]map[string]string {
		"oi":{
			"ola":"oi",
		},
	}

	fmt.Println(mp["oi"])

	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	for i,v := range tipos{
		fmt.Println(i, v)
	}

	for i,v := range mp{
		fmt.Println(i, v)
		for _,vv := range  v {
			fmt.Println(vv)
		}
	}

	fefe := Fernando{
		EhBonito:true,
		Pessoa: Pessoa{
			Nome:"Fernando",
			Idade:24,
		},
	}

	fmt.Println(fefe)
	fmt.Println(fefe.Pessoa.Nome)
	fefe.falar()
}
