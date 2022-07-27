package main

import (
	"fmt"

	"github.com/geofox-coding/first-go-app/calculator"
)

func main() {
	fmt.Println("my new go app knows the sum of 23 and 42 is ", calculator.Sum(23, 42))
	fmt.Println(calculator.Divide(17, 3))
}
