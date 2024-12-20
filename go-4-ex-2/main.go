package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(3, 4))  
	fmt.Println(computeHypotenuse(5, 12)) 
	fmt.Println(computeHypotenuse(8, 15)) 
}
