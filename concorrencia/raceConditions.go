package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup2 sync.WaitGroup
var result int
var result2 int
var m sync.Mutex

func init(){
	fmt.Println("confg")
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main() {

	fmt.Println(runtime.NumCPU())
	waitGroup2.Add(2)
	go runProcess2("P1", 20)
	go runProcess2("P2", 20)
	waitGroup2.Wait()
	fmt.Println("Final result: ",result)

	// Utilizando Mutex
	// Para descobrir race conditions : go run -race main.go
	// mutex == sincronized do java tipo um lock
	waitGroup2.Add(2)
	go runProcessComMutex("P1", 20)
	go runProcessComMutex("P2", 20)
	waitGroup2.Wait()
	fmt.Println("Final result2: ",result2)

}

func runProcess2(name string, total int){
	for i:=0;i<total;i++{
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, ">",i, "Partial result:",result)
	}
	waitGroup2.Done()
}

func runProcessComMutex(name string, total int){
	for i:=0;i<total;i++{
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result2++
		fmt.Println(name, ">",i, "Partial result2:",result)
		m.Unlock()
	}
	waitGroup2.Done()
}
