package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Rafi", filterName)

	// --> Or we can use like this:
	filter := filterName
	sayHelloWithFilter("Anjing", filter)
}
