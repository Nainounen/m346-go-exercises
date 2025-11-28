package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	dayOfBirth   byte
	monthOfBirth byte
	yearOfBirth  uint16
}
type Profile struct {
	ProfileName      FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		ProfileName: FullName{
			FirstName: "Nino",
			LastName:  "Meier",
		},
		Birth: BirthDate{
			dayOfBirth:   12,
			monthOfBirth: 4,
			yearOfBirth:  2008,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '\u2648',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
