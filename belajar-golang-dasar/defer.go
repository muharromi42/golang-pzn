package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	//defer pasti akan tereksekusi function nya terakhir sekelai walaupun function tersebut error dsb
	defer logging()
	fmt.Println("run application")
}

func main() {
	runApplication()
}
