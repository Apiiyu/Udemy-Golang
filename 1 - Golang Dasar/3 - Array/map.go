package main

import "fmt"

/*
	Map adalah tipe data lain yang berisikan kumpulan data yang sama,
	namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.

	Sederhananya, Map adalah tipe data kumpulan key-value, dimana kata kuncinya
	bersifat unik atau tidak boleh sama.
*/

func main() {
	person := map[string]string{
		"name":    "John",
		"age":     "20",
		"address": "Jl. Kebon Kacang",
	}

	person["occupation"] = "Software Engineer" // Menambahkan data ke map

	// --> Function in map
	fmt.Println(len(person))   // Mendapatkan panjang map
	fmt.Println(person["age"]) // Mendapatkan data pada map berdasarkan key
	delete(person, "address")  // Menghapus data pada map berdasarkan key
}
