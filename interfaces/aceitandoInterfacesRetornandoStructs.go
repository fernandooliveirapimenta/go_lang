package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	s := `{"Name":"Joe", "Age":18}`
	s2 := `{"Name":"Jane", "Age":21}`

	p, err := loadPerson(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	p2, err := loadPerson2(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	var str interface{} = "some string"
	v := str.(string)
	fmt.Println(strings.Title(v))

	fmt.Println()
	fmt.Println()
	fmt.Println()
	//var str2 interface{} = 49
	//v2 := str2.(string)
	//fmt.Println(strings.Title(v2))

	i := []interface{}{42, "The book club", true, time.Now()}
	typeExample(i)
}

func typeExample(i []interface{}){
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n",v)
		case string:
			fmt.Printf("%v is string\n",v)
		case bool:
			fmt.Printf("%v is bool\n",v)
		default:
			fmt.Printf("Unknown type %T",v)
		}
	}
}

func emptyDetails(s interface{}){
	fmt.Printf("(%v, %T)\n",s,s)
}

func loadPerson2(s string) (Person, error){
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p,nil
}

func loadPerson(r io.Reader) (Person, error){
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p,err
	}
	return p,nil
}