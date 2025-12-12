package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) string {
	D := b*b - 4*a*c

	if D < 0 {
		return fmt.Sprintf("no real solutions (D=%.5f)", D)
	}
	if D == 0 {
		x := -b / (2 * a)
		return fmt.Sprintf("one solution: x=%.5f (D=%.5f)", x, D)
	}
	sqrtD := math.Sqrt(D)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return fmt.Sprintf("two solutions: x1=%.5f, x2=%.5f (D=%.5f)", x1, x2, D)
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // two solutions: x1=-0.33333, x2=-1.00000 (D=4.00000)
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // one solution: x=-1.00000 (D=0.00000)
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // no real solutions (D=-8.00000)
}
