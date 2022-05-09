package main

import "fmt"

func main() {
	// --> Boolean Operation (&&, ||, !) (AND, OR, NOT)
	// --> Konsepnya seperti gerbang logika (AND, OR, NOT)
	var nilaiAkhir = 90
	var absensi int = 80

	lulusNilaiAkhir := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusNilaiAkhir && lulusAbsensi // --> AND True + True = True, Besides that it will be false
	fmt.Println(lulus)
}
