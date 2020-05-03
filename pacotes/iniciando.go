package main

//aliases
import (
	"fmt"
	m "github.com/fernandooliveirapimenta/go_lang/msg"
	"strings"
)

func main() {
	str := "found me"


	if strings.Contains(str, "found") {
		fmt.Println("value found in str")
	}

	m.Greeting("Reusando")

	a := m.AlunoF{
		Nome: "Fernando",
		Idade: 24}
	fmt.Println(a)

	a.Estudar()

}
