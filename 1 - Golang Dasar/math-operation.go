package main

import "fmt"

func main() {
	const panjangPersegi = 10
	const lebarPersegi int = 20

	luasPersegi := panjangPersegi * lebarPersegi // --> Operasi Perkalian
	fmt.Println(luasPersegi)

	// --> Augmented Assignments
	// --> +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=

	var nilaiIndonesia = 80
	var nilaiMatematika = 85
	nilaiIndonesia += nilaiMatematika
	fmt.Println(nilaiIndonesia)

	// --> Unary Operator
	// --> ++, --, +, -, !

	indexLooping := 0
	indexLooping++ // --> Increment (indexLooping = indexLooping + 1)
	fmt.Println(indexLooping)

	valuePositive := +1 // --> Unary Plus (Optional using +	sign, cause by default it is unary plus)
	valueNegative := -1
	fmt.Println(valuePositive, valueNegative)
}
