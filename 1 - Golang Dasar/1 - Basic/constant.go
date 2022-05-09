package main

import "fmt"

func main() {
	const pi float64 = 3.14              // --> Deklarasi constant/konstanta pi dengan tipe data float64
	const description = "Luas Lingkaran" // --> Deklarasi constant description dengan tipe data string tanpa menyebutkan tipe data
	const radius = 10                    // --> Constant tidak akan error jika tidak digunakan

	fmt.Println(description, pi)

	// Multiple Constant Declaration
	const ( // --> Cara penggunaannya hampir sama dengan Multiple Variable Declaration
		author    string = "Rafi"
		birthYear        = 2004
		isAlive          = true
	)

	fmt.Println(author, birthYear, isAlive)
}
