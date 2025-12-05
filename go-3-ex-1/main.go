package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	//used LLM for this formula
	zodiacNumber := int(p.Month)*100 + int(p.Day)
	if zodiacNumber >= 321 && zodiacNumber <= 420 {
		zodiacSign = Aries
	} else if zodiacNumber >= 421 && zodiacNumber <= 521 {
		zodiacSign = Taurus
	} else if zodiacNumber >= 522 && zodiacNumber <= 621 {
		zodiacSign = Gemini
	} else if zodiacNumber >= 622 && zodiacNumber <= 722 {
		zodiacSign = Cancer
	} else if zodiacNumber >= 723 && zodiacNumber <= 822 {
		zodiacSign = Leo
	} else if zodiacNumber >= 823 && zodiacNumber <= 922 {
		zodiacSign = Virgo
	} else if zodiacNumber >= 923 && zodiacNumber <= 1022 {
		zodiacSign = Libra
	} else if zodiacNumber >= 1023 && zodiacNumber <= 1122 {
		zodiacSign = Scorpius
	} else if zodiacNumber >= 1123 && zodiacNumber <= 1220 {
		zodiacSign = Sagittarius
	} else if zodiacNumber >= 1221 || zodiacNumber <= 119 {
		zodiacSign = Capricornus
	} else if zodiacNumber >= 120 && zodiacNumber <= 218 {
		zodiacSign = Aquarius
	} else if zodiacNumber >= 219 && zodiacNumber <= 320 {
		zodiacSign = Pisces
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)
}
