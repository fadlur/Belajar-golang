package main

import "fmt"

func main() {
	fmt.Println("Array adalah kumpulan data yang memiliki tipe data yang sama")
	fmt.Println("Array memiliki kapasitas yang ditentukan pada saat deklarasi")

	var names = [4]string{"John", "Doe", "Jane", "Doe"}
	fmt.Println(names)
	fmt.Println(names[0], names[1], names[2], names[3])

	fmt.Println("Atau bisa juga seperti ini")
	fmt.Println("Dideklarasikan dengan ukuran 4")
	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"
	fruits[3] = "Pineapple"

	fmt.Println(fruits)
	fmt.Println("Akan error kalau mengisi data lebih dari kapasitas array")
	fmt.Println("Semisal mau mendeklarasikan array tapi belum tau isine")
	fmt.Println("Maka bisa menggunakan ...")
	var numbers = [...]int{2, 3, 5, 7, 11}

	fmt.Println("data array: ", numbers)
	fmt.Println("jumlah data: ", len(numbers))

	fmt.Println("Array multidimensi")
	var numbers1 = [2][3]int{{3, 2, 3}, {2, 3, 4}}
	var numbers2 = [2][3]int{{3, 2, 3}, {2, 3, 4}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	fmt.Println("Perulangan Elemen Array menggunakan for")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	fmt.Println("Perulangan Elemen Array menggunakan for range")
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	fmt.Println("Alokasi Element array menggunakan make")
	var buah = make([]string, 2)
	buah[0] = "Apple"
	buah[1] = "Banana"

	fmt.Println(buah)
}
