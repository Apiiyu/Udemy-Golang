package main

import "fmt"

func withoutVarags(params []int) int {
	var total int
	for _, number := range params {
		total += number
	}

	return total
}

func sumAll(numbers ...int) int { //-> Variadic Function
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func validateVarags(name string, collectionNumbers ...int) {
	/*
		If we want to use variadic function, we must use ... in the function parameter.

		Note: Varags should be place at the last parameter.
	*/
	fmt.Println(name, collectionNumbers)
}

func main() {
	/*
		-- Variadic Function
		Variadic function adalah function yang memiliki parameter yang banyak.

		(Varags === Variable Arguments)

		1. Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs.
		2. Varags artinya datanya bisa menerima banyak nilai. (Anggap seperti array)

		-- Perbedaan parameter biasa dengan tipe data array
		1. Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function.
		2. Jika parameter menggunakan varags, kita bisa langsung mengirim datanya ke function. Jika lebih dari satu, cukup gunakan tanda koma untuk menandai.
	*/
	slice := []int{10, 22, 33, 44, 55}
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(slice...)) //-> Make slice to varags
}
