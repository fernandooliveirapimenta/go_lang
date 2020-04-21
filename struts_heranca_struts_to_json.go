package main

import (
	"encoding/json"
	"fmt"
)

type Car1 struct {
	Name string `json:"nome"`
	Year int `json:"-"`
	color string
	Cod string

}

//Heranca basta add o tipo na struct
type SuperCar struct {
    Car1
	CanFly bool
    Name string
}

//Estou attachando esse funcao na struct

func(c Car1) info() string{
	return fmt.Sprintf("Car: %v \n Year: %v \n Color: %v\n",
		c.Name, c.Year, c.color)
}

func main() {

	car1 := Car1{"Corolla", 2017, "Black", "2020"}
	fmt.Println(car1)

	sCar := SuperCar{
		Car1: Car1{
			"Fusca",
			2030,
			"Blue",
			"2020",
		},
		CanFly:true,
		Name: "Elantra",
	}

	fmt.Println(sCar)
	fmt.Println(sCar.Name)
	fmt.Println(sCar.Car1.Name)


	//Json

	cars := Car1{"Gol", 2017, "Yellow", "2020"}
	// so exporta atributos public
	re , _ := json.Marshal(cars) // retorna binario do json
	fmt.Println(string(re))

	carJson := []byte(`{"nome":"Gol", "Year":2017, "Cod":"2020202"}`)
	fmt.Println(string(carJson))

	var cc Car1
	json.Unmarshal(carJson, &cc)
	fmt.Println(cc)
	fmt.Println(cc.color)
	fmt.Println(cc.Cod)

}
