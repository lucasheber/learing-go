package main

import (
	"fmt"
	"math"
)

func main() {
	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	if x := math.Sqrt(4); x < 10 {
		fmt.Println("sqrt")
	}
}
