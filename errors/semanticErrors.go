package main

import (
	"errors"
	"fmt"
	"strconv"
)


type errorString struct {
	s string
}
func main() {

	km := 2

	if km > 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Going to walk today")
	}

	lancarErro()
	//panic(lancarErro())


	// Error interface type
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n",s,s)
	} else {
		fmt.Println(err)
	}
	v = "s2"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(s, err)
	}

	// meu tipo

	//es := errors.errorString{}
	//es.s = "slacker"
	//fmt.Println(es)

	errBadData := errors.New("Some bad data")
	fmt.Printf("ErrBadDAta type: %T", errBadData)

	errors.New("Http")


}

func New(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string{
	return e.s
}
func lancarErro() error {
	return fmt.Errorf("Erro sim")
}
