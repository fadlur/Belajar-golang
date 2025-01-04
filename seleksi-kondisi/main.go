package main

import "fmt"

func main() {
	// sudah pahamlah soal if else
	// variable temporary yang belum sering pake
	// kenapa pake temporary variable:
	// 1. scope variable jelas, hanya digugnakan di if else
	// 2. kode menjadi rapi
	// 3. ketika nilai variable didapat dari komputasi, perhitungan tidak perlu dilakukan di dalam blok masing-masing
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// dari contoh di atas, percent didapat dari hasil perhitungan, dan hanya digunakan di deretan blok seleksi kondisi itu saja
	// yang mencakup blog if, else if, dan else

	// switch case
	var point2 = 6
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// pemanfaatan case untuk banyak kondisi
	var point3 = 6
	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{ // lucunya di golang, kita bisa buat blok di dalam case
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// switch case dengan gaya if else
	var point4 = 6
	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// penggunaan fallthrough di switch case
	// di golang switch case akan berhenti ketika sudah menemukan case yang sesuai
	// jika kita ingin melanjutkan ke case selanjutnya, kita bisa menggunakan fallthrough
	var point5 = 6
	switch {
	case point5 == 8:
		fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("awesome")
		fallthrough
	case point5 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// jadi kalau ada fallthrough, maka case selanjutnya akan dijalankan dan ketika memenuhi case selanjutnya, maka akan berhenti

	// switch case bersarang
	var point6 = 6
	if point6 > 7 {
		switch point6 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point6 == 5 {
			fmt.Println("awesome!")
		} else {
			fmt.Println("not bad!")
		}
	}

}
