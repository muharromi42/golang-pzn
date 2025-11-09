package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var romi Customer
	romi.Name = "romi"
	romi.Address = "japan"
	romi.Age = 18

	fmt.Println(romi)
	fmt.Println(romi.Name)
	fmt.Println(romi.Address)
	fmt.Println(romi.Age)

	// bisa juga membuat struct literal atau langsung mendifinisikan isinya
	albert := Customer{
		Name:    "albert",
		Address: "japan",
		Age:     12,
	}
	fmt.Println(albert)

	kevin := Customer{"kevin", "japan", 20}
	fmt.Println(kevin)
}
