package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go runProcess1("P1", 20)
	go runProcess1("P2", 20)

	//go routines mais leve que threads em java
	// paralelismo

	var s string
	fmt.Scanln(&s)

}

func runProcess1(name string, total int){
	for i:=0;i<total;i++{
		fmt.Println(name, ">",i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
