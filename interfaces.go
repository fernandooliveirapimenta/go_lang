package main

import "fmt"

type car struct {
	name string
}
func (c car) start() string {
	return "The car" + c.name + " har been started"
}

type motorcycle struct {
	name string
}
func (c motorcycle) start() string {
	return "The motorcycle" + c.name + " har been started"
}

//Quero fazer uma funcao generica ou seja com a mesma assinatura
type vehicle interface {
	start() string
}

func startVehicle(v vehicle) string {
	return v.start()
}

type Names []interface{}

func (n *Names) load(){
	*n = Names{
		"Wslay",
		"Sara",
		"Davi",
		5.5,
	}
}

func (n Names) printNames(){
	fmt.Println(n)
}

func main() {

	c := car{"Gol"}
	fmt.Println(startVehicle(c))

	fmt.Println()

	m := motorcycle{"Gol"}
	fmt.Println(startVehicle(m))


	//Interface vazia

	var names Names
	names.load()
	names.printNames()

}