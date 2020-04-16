package main

import (
	"fmt"
	"math/rand"
)

func main() {


	// Using the for i Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


	// Looping over arrays and slices
	names := []string{"Jim", "Jane", "Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}


	// Loop over a map
	config := map[string]string{
		"debug": "1",
		"logLevel": "warn",
		"version" : "1.2.1",
	}

	f :=  config["logLevel"];
	fmt.Println(f)
	for key, value := range config {
		fmt.Println(key, "=", value)
	}


	// Break and continue
	for {
		r := rand.Intn(8)
		if r % 3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r % 2 == 0 {
			fmt.Println("Stop")
			break
		}

		fmt.Println("random: ", r)

	}




}
