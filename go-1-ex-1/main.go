package main

import "fmt"

func main() {
	var firstName string = "Max"
	var lastName string = "Muster"
	var dayOfBirth int = 22
	var monthOfBirth int = 11
	var yearOfBirth int = 2000
	var numberOfSiblings int = 2
	var heightInMeters float64 = 1.75
	var zodiacSign rune = '\u2650' 

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
