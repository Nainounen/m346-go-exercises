package main

import "fmt"

func main() {
	modules := map[uint]string{
		104: "Ein modul das ich nicht kenne",
		117: "Ein weiteres Modul das ich nicht kenne",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104)

	modules[320] = "c#"

	modules[117] = "Ich kenne es doch"
	fmt.Println(modules)
}
