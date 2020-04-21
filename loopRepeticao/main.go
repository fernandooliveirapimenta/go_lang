package main

import "fmt"

func main() {

	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	fmt.Println()

	x := 0
	for x < 10 {
		fmt.Println(x)
		x ++
	}


	fmt.Println()
	stop := 0
	for {
		fmt.Println(stop)
		if stop == 10 {
			break
		}

		stop ++
	}

}
