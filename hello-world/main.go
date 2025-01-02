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
}
