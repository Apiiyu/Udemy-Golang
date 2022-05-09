package main

import "fmt"

func main() {
	// --> Array
	// --> Array adalah sebuah variabel yang bisa menampung banyak nilai
	// --> Array bisa menampung data berupa string, integer, boolean, ataupun float

	// --> First way to declare an array
	var arrayNama [3]string
	arrayNama[0] = "Rizki"
	arrayNama[1] = "Rafi"
	arrayNama[2] = "Budi"

	fmt.Println(arrayNama[0])
	fmt.Println(arrayNama[1])
	fmt.Println(arrayNama[2])

	// --> Second way to declare an array
	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)

	fmt.Println(len(arrayNama)) // --> Length of array
	fmt.Println(cap(arrayNama)) // --> Capacity of array
}
