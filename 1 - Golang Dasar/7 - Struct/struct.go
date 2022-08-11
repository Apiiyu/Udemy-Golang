package main

import "fmt"

type User struct {
	Name     string
	Username string
	Email    string
	Passwrod string
}

func (users User) sayHello(age int) { // --> Struct method
	fmt.Println(users.Name, age)
}

func main() {
	/*
		-- Struct
		1. Struct adalah sebuah template data yang menggabungkan nol atau lebih tipe data lainnya dalam suatu kesatuan.
		2. Struct biasanya representasi data dalam program aplikasi yang kita buat.
		3. Data di struct disimpan didalam field

		   -- Conclusion : Struct adalah kumpulan dari field
	*/
	user := User{
		Name:     "John Doe",
		Username: "jdoe",
		Email:    "john@gmail.com",
		Passwrod: "12345",
	}

	fmt.Println(user.Name)
	user.sayHello(30)
}
