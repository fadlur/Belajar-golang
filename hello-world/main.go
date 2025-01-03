package main

import "fmt"

func main() {
	fmt.Println("Hello, World!", "how", "are", "you?")

	var firstName string = "john"

	var lastName string = "wick"

	fmt.Printf("Halo %s %s !\n", firstName, lastName)

	/*
		Deklarasi variabel bisa juga menggunakan new. ini digunakan untuk membuat variable pointer dengan tipe data tertentu.
	*/

	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)

	/*
		Deklarasi variable bisa juga menggunakan make. ini digunakan untuk:
		- channel
		- slice
		- map
	*/

	// variable

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// variable decimal cukup tau 2 jenis, yaitu float32 dan float64
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// variable boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// variable string
	var message string = "Halo"
	fmt.Printf("message: %s\n", message)

	var pesan = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(pesan)

	// konstanta menggunakan multiple
	const (
		namaDepan = "john"
		namaBelakang
	)

	// kalau konstanta kedua tidak diberi nilai dan type, maka otomatis nilainya seperti atasnya

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
