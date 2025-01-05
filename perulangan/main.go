package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// perulangan biasanya
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Perulangan hanya dengan kondisi (gaya while)
	var i1 = 0

	for i1 < 5 {
		fmt.Println("Angka", i1)
		i1++
	}

	// Perulangan tanpa argumen, dihentikan dengan break
	var i2 = 0

	for {
		fmt.Println("Angka", i2)
		i2++
		if i2 == 5 {
			break
		}
	}

	// Perulangan menggunakan for - range
	// biasa digunakan untuk looping array atau slice
	var xs = "123"
	for a, v := range xs {
		fmt.Println("index", a, "=", v)
	}

	var ys = [5]int{1, 2, 3, 4, 5}
	for _, v := range ys {
		fmt.Println("Value: ", v)
	}

	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3}
	for k, v := range kvs {
		fmt.Println("Key: ", string(k), "Value: ", v)
	}

	// k dan atau v diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// bisa juga cukup menentukan nilai numeriknya
	for i := range 5 {
		fmt.Println(i)
	}

	// Penggunaan continue dan break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Perulangan bersarang dengan label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	// masih bingung dengan label, coba perhatikan contoh ini
outerLoop2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop2
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
