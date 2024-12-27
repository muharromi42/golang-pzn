package main

import "fmt"

func main() {
	names := [...]string{
		"romi",
		"ardhana",
		"albert",
		"romay",
		"kevin",
		"stuart",
	}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// untuk mengecek panjang array dri slice
	fmt.Println(len(slice1))

	// untuk mengecek kapasitas dari array
	fmt.Println(cap(slice1))

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(days)

	daySlice1 := days[5:] // sabtu minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "Sabtu lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("===================================================")

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "romi"
	newSlice[1] = "romay"
	// newSlice[2] = "romay" // bakal error harus menggunakan append untuk menambahkan data di array, karna di awal sudah di inisiasikan panjang array nya hanya 2

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Muharrom")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "albert"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("===================================================")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("===================================================")

	// hati hati saat membuat array karna hanya beda ...
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}