package main

import (
	"errors"
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) ([]float64, error) {
	if a == 0 {
		return nil, errors.New("ungültiger Wert für a: Die Gleichung ist nicht quadratisch")
	}

	discriminant := math.Pow(b, 2) - 4*a*c

	switch {
	case discriminant > 0:
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{root1, root2}, nil
	case discriminant == 0:
		root := -b / (2 * a)
		return []float64{root}, nil
	default:
		return nil, errors.New("keine reellen Lösungen")
	}
}


func main() {

	roots1, err1 := computeQuadraticFormula(3, 4, 1)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Lösungen:", roots1) 
	}

	roots2, err2 := computeQuadraticFormula(2, 4, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Lösungen:", roots2) 
	}

	roots3, err3 := computeQuadraticFormula(3, 4, 2)
	if err3 != nil {
		fmt.Println(err3) 
	} else {
		fmt.Println("lösungen", roots3)
	}
}

