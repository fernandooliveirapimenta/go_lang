package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(5) + 1
	strars := strings.Repeat("*", r)
	fmt.Println(strars)

}