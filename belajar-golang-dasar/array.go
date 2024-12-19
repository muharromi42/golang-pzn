package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "romi"
	names[1] = "putra"
	names[2] = "ardhana"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	// bisa langsung inisialisasi array
	var values = [3]int{
		1,
		2,
		3,
	}


	var values2 = [...]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(values)
	fmt.Println(values2)
	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(values2))
}