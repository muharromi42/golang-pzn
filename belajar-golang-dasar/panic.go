package main

import "fmt"

func endApp() {
	fmt.Println("end App")
}

func runApp(error bool) {
	//panic untuk menghentikan app secara menyeluruh tetapi defer tetap berjalan
	defer endApp()
	if error {
		panic("upss error")
	}
}

func main() {
	runApp(true)
}
