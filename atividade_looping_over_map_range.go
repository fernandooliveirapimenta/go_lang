package main

import "fmt"

func main() {
	// Looping over map data using range

	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up": 4,
	}

	bigCount := 0
	bestWorld := ""
	for key, value := range words {
		if value > bigCount {
			bigCount = value
			bestWorld = key
		}

	}

	fmt.Println("Most popular word: ", bestWorld)
	fmt.Println("Most a count of: ", bigCount)
}
