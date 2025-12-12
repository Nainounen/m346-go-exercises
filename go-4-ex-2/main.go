package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(s ShortSides) float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(ShortSides{2.4, 5.7}))   //6.18465843842649
	fmt.Println(computeHypotenuse(ShortSides{25.2, 12.6})) //28.17445651649735
	fmt.Println(computeHypotenuse(ShortSides{21.4, 9.0}))  //23.21551205552012
}
