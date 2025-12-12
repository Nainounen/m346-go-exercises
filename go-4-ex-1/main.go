package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float32) (float32, error) {
	if gotPoints == 0 {
		return 0.0, errors.New("Cant calculate your grade, you didnt get any Points")
	}
	return gotPoints/maxPoints*5 + 1, nil
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125 <nil>
	fmt.Println(computeGrade(25.5, 28.0)) // 5.553571 <nil>
	fmt.Println(computeGrade(0.0, 28.0))  // 0 Cant calculate your grade, you didnt get any Points
}
