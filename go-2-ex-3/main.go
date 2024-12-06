package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Programmieren Grundlagen",
		117: "Algorithmen und Datenstrukturen",
		346: "Cloud-Lösungen konzipieren",
	}

	delete(modules, 117)
	modules[200] = "Projektmanagement"
	modules[346] = "Cloud-Infrastruktur"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117]) 
	fmt.Println("Modul 346:", modules[346])
	fmt.Println(modules)
}
