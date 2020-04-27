package main

import (
	"fmt"
	"os"
	"time"
)

type Speaker interface {
	Speak(message string) string
	Greet() string
}

type Reader interface {
	Read(p []byte)(n int, err error)
}

type cat struct {
	name string
	age int
}

func (c cat) Speak() string{
	return "Purr Meow"
}

func (c cat) Greeting() {
	fmt.Println("Meow,Meow !!!! mmmeeeeoooowwwwww")
}

func (c cat) String() string{
	return fmt.Sprintf("%v (%v years old)",c.name, c.age)
}

type FileInfoFernando interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}
}


func main() {
	c := cat{name: "Oreo", age:9}
	fmt.Println(c.Speak())
	fmt.Println(c)
	//os.FileInfo()
}

