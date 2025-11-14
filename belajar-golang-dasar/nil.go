package main

import "fmt"

// nil data kosong
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
