package main

import "fmt"

func main() {
	/*
		Switch adalah perintah yang digunakan untuk mengkondisikan kondisi tertentu.
		Switch expression sangat sederhana dibandingkan if expression.
		Switch expression memiliki 2 bagian, yaitu switch expression dan case expression.

		Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam
		satu variable.
	*/

	const username = "Rafi Khoirulloh"

	switch username {
	case "Rafi":
		fmt.Println("Hello Rafi")
	case "Rafi Khoirulloh":
		fmt.Println("Hello Rafi Khoirulloh")
	default:
		fmt.Println("Hello John Doe")
	}

	// --> Switch short hand statement
	switch lengthUsername := len(username); lengthUsername > 5 {
	case true:
		fmt.Println("Username is more than 5")
	case false:
		fmt.Println("Username is less than 5")
	default:
		fmt.Println("Username is 5")
	}
}
