package main

import "fmt"

func main() {
	var name string = "Rafi" // --> Deklarasi variabel name dengan tipe data string
	fmt.Println(name)        // --> Di Golang setiap variable harus digunakan, jika tidak maka akan error

	var age = 23 // --> Deklarasi variabel age dengan tipe data integer tanpa menyebutkan tipe data
	fmt.Println(age)

	address := "Jl. Raya" // --> Deklarasi variabel tanpa menggunakan kata kunci var
	fmt.Println(address)

	// Multiple Variable Declaration
	var name1, name2, name3 string = "Rafi1", "Rafi2", "Rafi3" // --> Cara ke-1 untuk membuat Multiple Variable
	fmt.Println(name1, name2, name3)

	var ( // --> Cara ke-2 untuk membuat Multiple Variable
		name4  = "Rafi4"
		status = true
	)
	fmt.Println(name4, status)

	age1, age2, age3 := 23, 24, 25 // --> Cara ke-3 untuk membuat Multiple Variable
	fmt.Println(age1, age2, age3)
}
