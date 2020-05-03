package main

import "github.com/fernandooliveirapimenta/go_lang/Exercise8.01/shape"

func main(){
	t := shape.Triangle {
		Base: 15.5,
		Height: 20.1,
	}

	r := shape.Rectangle{
		20, 10,
	}

	s := shape.Square{
		10,
	}

	shape.PrintShapeDetails(t, r, s)
}

