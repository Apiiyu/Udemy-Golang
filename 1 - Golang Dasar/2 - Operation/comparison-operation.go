package main

import "fmt"

func main() {
	// --> Comparison Operator (==, !=, <, >, <=, >=)
	// --> Comparsion Operator beguna untuk membandingkan 2 buah nilai
	// --> Comparison Operator akan menghasilkan value boolean (true or false)

	const pi float64 = 3.14
	const secondPi = 22 / 7
	comparisonPi := pi == secondPi
	fmt.Println(comparisonPi)

	var valueIndonesia int = 100
	valueEnglish := 90
	fmt.Println(valueEnglish != valueIndonesia)
	fmt.Println(valueEnglish < valueIndonesia)
	fmt.Println(valueEnglish > valueIndonesia)
	fmt.Println(valueEnglish <= valueIndonesia)
	fmt.Println(valueEnglish >= valueIndonesia)
}
