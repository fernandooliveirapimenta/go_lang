package main

import "fmt"

// forma como o go trabalha com orientacao a objetos
// é tipo uma classe do Java

type SuperNumber int

type PP SuperNumber

type CarName string
type CarYear int
type Color string
type Date int

type Car struct {
	Name string
	Year int
	Color string
}

//Estou attachando esse funcao na struct
func (c *Car) ligar() {
	c.Year = 50
	fmt.Println("Ligando: ",c.Name)
	fmt.Println("Ligando: ",c.Color)
	fmt.Println("Ligando: ",c.Year)
}

func main() {

	x := []SuperNumber{3,3}
	fmt.Println(x)

	car1 := Car{"Corolla", 2017, "white"}
	car2 := Car{"Bmw 320i", 2017, "Red"}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.Name)
	fmt.Println(car2.Color)
	car1.ligar()
	fmt.Println(car1.Year)

}
