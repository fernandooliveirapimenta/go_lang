package main

import (
	"fmt"
	"time"
)

const (
	n1 = iota
	n2
	n3
)

func main(){

	fmt.Println(n1,n2,n3)

	var (
		nome string = "fernando"
		idade int
		dtNascimento time.Time
	)

	config := map[string]string {
		"env":"dev",
		"env2":"dev",
		"env3":"dev",
		"env4":"dev",
	}

	idade, dtNascimento = 24, time.Date(1995, 10, 30, 0, 0,0,0, time.Local)

	fmt.Printf("Nome %v, idade %v, dtNascimento %v \n",nome, idade, dtNascimento)


	for key, v := range config {
		fmt.Printf("key: %s, value: %s\n", key, v)
	}

	for i := 0; i < 10; i++ {

		fmt.Printf(string(i * (i +1)))

	}

	i := 0
	for {
		if i == 3 {
			break
		} else {
			i ++
			fmt.Println("continue")
			continue
		}
	}

	segredos := []string {"ola", "oi"}

	fmt.Println(segredos)
	segredos = append(segredos, "gogo", "hehe")
	fmt.Println(segredos)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(segredos[:2])
	fmt.Println(segredos[2+1:])
	segredos = append(segredos[:2], segredos[2+1:]...)
	fmt.Println(segredos)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	sm := map[string]string{
		"2":"1",
		"3":"1",
		"4":"1",
		"5":"1",
		"6":"1",
	}
	fmt.Println(sm)
	sm["7"] = "7"
	fmt.Println(sm)

	delete(sm, "4")
	fmt.Println(sm)

	fmt.Println()
	fmt.Println()
	fmt.Println()

}
