package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func onSayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func onReturnValue(firstNumber int, secondNumber int) int { //-> Function with return value
	return firstNumber + secondNumber
}

func multipleValue() (string, int) { //-> Function with multiple return value
	return "Hello", 10
}

func getFullName() (string, string, string) {
	return "Rafi", "Khoirulloh", "BTS.id"
}

func main() {
	sayHello()
	onSayHello("John", "Doe")
	onSayHello("Rafi", "Khoirulloh")
	result := onReturnValue(10, 20)
	firstValue, secondValue := multipleValue()
	firstName, lastName, _ := getFullName()
	fmt.Println(firstValue, secondValue)
	fmt.Println(result)
	fmt.Println(firstName, lastName)
}
