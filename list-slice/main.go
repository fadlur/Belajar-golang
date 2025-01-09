package main

import "fmt"

func main() {
	fmt.Println("Slice adalah potongan dari array")
	fmt.Println("Slice merupakan reference array, jadi ketika ada perubahan maka akan berdampak pada slice lain yang memiliki alamat memori yang sama")
	// Inisialisasi Slice
	var buah = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(buah)
	// ambil index pertama
	var aBuah = buah[0]
	fmt.Println(aBuah)
	// perbedaaan slice dan array adalah saat deklarasi slice tidak perlu menentukan jumlah elemen
	// slice := []string{"apple", "grape", "banana", "melon"}
	// array := [4]string{"apple", "grape", "banana", "melon"}

	// Hubungan antara slice dan array
	// Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut
	// Jika kita mengubah nilai elemen pada array, maka slice yang mengacu pada array tersebut juga akan berubah

	// Slice merupakan tipe data reference
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bbFruits[0] = "pinnaple"
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	// Fungsi len() dan cap()
	// len() digunakan untuk menghitung jumlah elemen pada slice
	// cap() digunakan untuk menghitung kapasitas maksimal slice
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4, ini 4 karena mulainya dari index 1, jadi 0 ikut dihitung

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// Fungsi append()
	// append() digunakan untuk menambahkan elemen pada slice
	// append() akan membuat slice baru jika kapasitas tidak cukup
	var animals = []string{"cat", "dog", "fish"}
	fmt.Println(animals)
	var new_animals = append(animals, "tiger")
	fmt.Println(new_animals)
	// Fungsi copy()
	// copy() digunakan untuk menyalin elemen dari slice ke slice lain
	var cFruits = []string{"apple", "grape", "banana"}
	var dFruits = make([]string, 2)
	copy(dFruits, cFruits)
	fmt.Println(cFruits)

	// Pengaksesan elemen slice dengan 3 index
	// slice[low:high:max]
	// low = index awal
	// high = index akhir
	// max = kapasitas maksimal slice
	var eFruits = []string{"apple", "grape", "banana", "melon"}
	var new_eFruits = eFruits[0:2:3]
	fmt.Println(new_eFruits)
}
