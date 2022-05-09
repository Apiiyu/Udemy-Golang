package main

import "fmt"

func main() {
	// --> Slice
	// Slice adalah potongan data array. Slice mirip seperti array namun slice ukurannya bisa berubah.
	// Keywords in slice (Pointer, Length, Capacity)
	// Pointer: menunjukan alamat data pada slice
	// Length: menunjukan panjang data pada slice
	// Capacity: menunjukan kapasitas data pada slice, (Length <= Capacity)

	// --> Ways to create slice
	// 1. array[low:high] (low dan high adalah index) - Artinya ialah membuat slice dari array dari index low - high
	// 2. array[low:] (low adalah index) - Artinya ialah membuat slice dari array dari index low - akhir array
	// 3. array[:high] (high adalah index) - Artinya ialah membuat slice dari array dari awal array - index high
	// 4. array[:] (Artinya ialah membuat slice dari array dari awal array - akhir array)

	// --> Function in slice
	// 1. append(slice, data) - Menambahkan data ke slice
	// 2. copy(slice, data) - Menyalin data ke slice
	// 3. len(slice) - Mendapatkan panjang slice
	// 4. cap(slice) - Mendapatkan kapasitas slice
	// 5. make([]TypeData, length, capacity) - Membuat slice dengan tipe data TypeData, panjangnya length, dan kapasitasnya capacity
	// 6. delete(slice, index) - Menghapus data pada slice pada index yang diberikan

	var months = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var firstSlice = months[4:7]
	var secondSlice = months[10:]

	fmt.Println(cap(firstSlice))
	fmt.Println(len(firstSlice))

	firstSlice[0] = "Change data array"
	fmt.Println(months)
	fmt.Println(secondSlice)
}
