package main

import "fmt"

func main() {
	items := map[string]map[string]float32{
		"cake": {"cost": 0.99, "tax": 7.5},
		"milk": {"cost": 2.75, "tax": 1.5},
		"butter": {"cost": 0.87, "tax": 2},
	}

	var taxTotal float32 = 0.0;
	for _, v := range items {

		tax := (v["cost"] * v["tax"]) / 100
		taxTotal += tax

	}

	fmt.Println("Stales Tax Total:",taxTotal)


	//sum := 0.99 + 2.75 + 0.87
	//
	//a1 := (0.99 * 7.5) / 100
	//a2 := (2.75 * 1.5) / 100
	//a3 := (0.87 * 2) / 100
	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(a3)
	//fmt.Println(a1+a2+a3)
	//fmt.Println((10 * 50) / 100 )
	//fmt.Println(sum)
}
