package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println(m)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	fmt.Println(m["c"])

	_,exists := m["c"]
	fmt.Println("existe c:",exists)
	_,existsb := m["b"]
	fmt.Println("existe b:",existsb)

	value, _ := m["a"]
	fmt.Println("valor:",value,)

	var x = map[string]int{
		"x": 5,
		"y": 10,
	}
	fmt.Println(x)

	if v, e := m["a"]; e {

		fmt.Println(v)

	} else {

		fmt.Println("nao exists")
	}

}

