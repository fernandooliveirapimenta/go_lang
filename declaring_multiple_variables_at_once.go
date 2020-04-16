package main

import (
	"fmt"
	"time"
)

var (
	debug bool = false
	logLevel string = "info"
	startUpTime time.Time = time.Now()
)

// sem tipo inferido

var (
	calvo bool
	corCabelo = "Loiro"
	anoAtual = time.Now().Year()
)


func main(){
	debug = true
	fmt.Println(debug, logLevel, startUpTime)

	fmt.Println(calvo, corCabelo, anoAtual)

}
