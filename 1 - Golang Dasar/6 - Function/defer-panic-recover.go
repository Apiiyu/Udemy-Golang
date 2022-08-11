package main

import "fmt"

func logger() {
	fmt.Println("Function already to execute")
}

func withoutDefer(value int) {
	defer logger() //--> Should will be executed even if there is an error.
	fmt.Println("Without Defer")
	result := 10 / value
	fmt.Println("result:", result)
	logger() //--> Should be not executed because of error.
}

func onDestroy() {
	errorMessage := recover()
	if errorMessage != nil {
		fmt.Println("error message:", errorMessage)
	}

	fmt.Println("Application is destroyed")
}

func runApplication(error bool) {
	defer onDestroy() //--> Defer function will be executed after run application success to executed.
	if error {
		panic("Application is ERROR!")
	}
	fmt.Println("Application is running")
}

func main() {
	/*
		-- Defer
		1. Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi.
		  --> Conclusion : Defer function ialah function yang bisa kita atur kapan akan dieksekusi.

		2. Defer function akan selalu dieksekusi walaupun terjadi error di function yang bersangkutan.

		-- Panic
		1. Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
		2. Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan.
		3. Saat Panic function dipanggil, program akan langsung dihentikan namun defer function tetap akan dieksekusi.

		-- Recover
		1. Rec over function adalah function yang bisa kita gunakan untuk menangani error pada saat program kita berjalan.
		2. Recover function juga bisa kita gunakan untuk menangkap data panic.
		3. Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.
	*/
	runApplication(true)
	// withoutDefer(0)
}
