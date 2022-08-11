package main

import "fmt"

/*
	Di Golang, kita tidak boleh menggunakan nama function yang sama didalam sebuah package. --> Important!
*/

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rafi"
	middleName = "Khoirulloh"
	lastName = "BTS.id"
	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
