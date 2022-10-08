package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	if condition := 3 == 3; condition {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	seedvalue := rand.Int63n(1000)
	fmt.Println(seedvalue)
}
