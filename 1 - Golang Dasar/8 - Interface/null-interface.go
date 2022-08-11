package main

import "fmt"

func Ups(params interface{}) interface{} {
	if params == 1 {
		return 1
	} else if params == true {
		return true
	} else {
		return "Ups" // --> It will be executed if function ups have parameter, even it integer/boolean or string.
	}
}

func main() {
	/*
		-- Null Interface
		1. Null interface adalah interface yang tidak memiliki nilai.
		2. Null interface bisa digunakan untuk mengatur nilai default.
		3. Null interface ialah interface yang tidak memiliki deklrasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya.
	*/
	nullInterface := Ups("asd")
	fmt.Println(nullInterface)
}
