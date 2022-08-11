package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is blacklisted")
	} else {
		fmt.Println("User", name, "is not blacklisted")
	}
}

func main() {
	blacklist := func(name string) bool { //-> Anonymous Function
		return name != "Rafi"
	}

	registerUser("Rafi", blacklist)
	registerUser("Anjing", func(name string) bool { //-> Anonymous Function
		return name == "Anjing"
	})
}
