package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Enter Value: ")
	var inputobj string
	fmt.Scan(&inputobj)
	fmt.Printf("Hello %v! Current Time is %v", inputobj, time.Now())
}
