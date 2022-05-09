package main

import "fmt"

func main() {
	// --> Type declaration berguna untuk membuat ulang tipe data yang baru dari tipe data yang sudah ada
	// --> Seperti membuat alias untuk sebuah tipe data
	type isMarried bool // --> Deklarasi tipe data isMarried dengan tipe data bool
	var married isMarried = true

	fmt.Println(married)
}
