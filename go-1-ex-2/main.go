package main

import "fmt"

func main() {
	const mileInKM = 1.60934
	var meilen = 55.3
	var kilometer = meilen * mileInKM
	fmt.Printf("%.2f Meilen = %.2f Kilometer\n", meilen, kilometer)

	var fahrenheit = 99.9
	var celsius = (fahrenheit - 32.0) * 5 / 9
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)

	const marathonInKM = 42.195
	var marathonInMeilen = marathonInKM / mileInKM
	fmt.Printf("Ein Marathon ist %.2f Kilometer = %.2f Meilen lang\n", marathonInKM, marathonInMeilen)

	var siedepunktWasserCelsius = 100.0
	var siedepunktWasserFahrenheit = siedepunktWasserCelsius*9/5 + 32.0
	fmt.Printf("Wasser siedet bei %.2f°C = %.2f°F\n", siedepunktWasserCelsius, siedepunktWasserFahrenheit)
}
