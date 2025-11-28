package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class map[byte]Student

func main() {
	class1 := Class{
		1: {FirstName: "Joe", LastName: "Doe"},
		2: {FirstName: "Jay", LastName: "Day"},
		3: {FirstName: "Wef", LastName: "EwfwefDay"},
	}

	class2 := Class{
		1: {FirstName: "Jim", LastName: "Jam"},
		2: {FirstName: "Jam", LastName: "Bam"},
		3: {FirstName: "Jan", LastName: "Beam"},
	}

	modules := map[uint][]Class{
		346: {class1, class2},
		320: {class1},
		114: {class2},
	}

	fmt.Println(class1)
	fmt.Println(class2)
	fmt.Println(modules)
}
