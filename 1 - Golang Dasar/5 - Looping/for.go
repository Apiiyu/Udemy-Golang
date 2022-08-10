package main

import "fmt"

func main() {
	/*
		Perulangan di Golang hanya ada 1 yaitu for.

		For adalah perintah yang digunakan untuk melakukan looping.
		For terdiri dari 3 bagian yaitu
		1. init statement - inisialisasi looping
		2. condition - kondisi looping
		3. post statement - perintah setelah looping
	*/

	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}

	/*
		For Range
		For range digunakan untuk melakukan looping pada slice, map, atau array.
	*/
	var months = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for index := 0; index < len(months); index++ {
		fmt.Println(months[index])
	}

	for index, month := range months {
		fmt.Println(index, month)
	}
}
