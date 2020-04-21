package main

import (
	"fmt"
	"github.com/fernandooliveirapimenta/go_lang/criandoPacotesEVisibilidade/car"
)

func main() {

	carr := car.Car{"Gol", "Black", "dd"}
	fmt.Println(carr)
	fmt.Println(carr.Color)
	fmt.Println("main")
}
