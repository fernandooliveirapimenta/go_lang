package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name string
	age int
	isMarried bool
}

type cat struct {

}
func (c cat) Speak() string{
	return "Cat speak "
}

type dog struct {

}

func (d dog) Speak() string{
	return "Dog speak"
}

func (p person) String() string{
	return fmt.Sprintf("%v (%v years old).\nMarried status: %v ", p.name, p.age, p.isMarried)
}

func (p person) Speak() string{
	return "Hi my name is: " + p.name
}



func main() {
	p := person{name:"Cailyn", age:44, isMarried:false}
	fmt.Println(p.Speak())
	fmt.Println(p)

	// Duck Typing
	chatter(p)

	// Polymorphism
	c:= cat{}
	d:= dog{}

	saySomething(p,c,d)
}

func saySomething(say ...Speaker){
	for _, s := range say{
		fmt.Println(s.Speak())
	}
}

func chatter(s Speaker){
	fmt.Println(s.Speak())
}
