package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)

	//closure kemampuan function berinteraksi dengan data-data di sekitarnya
	//	contoh program di atas function bisa mempengaruhi isi dari variabel counter yang padahal var nya ada di luar blok function
}
