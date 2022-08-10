package main

import "fmt"

func main() {
	/*
		Break adalah perintah yang digunakan untuk menghentikan perulangan.
		Continue adalah perintah yang digunakan untuk melanjutkan perulangan.
		Break dan continue bisa digunakan pada perulangan for, while, dan switch.
	*/
	for index := 1; index <= 10; index++ {
		if index == 5 {
			fmt.Println("Break")
			break
		}
		fmt.Println(index)
	}

	for index := 1; index <= 10; index++ {
		if index%2 == 0 {
			fmt.Println("Continue")
			continue
		}
		fmt.Println(index)
	}
}
