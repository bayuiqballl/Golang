package main

import "fmt"

// mencari luas lingkaran
func main() {
	var phi float32 = 3.14
	var r float32 = 12 //jari jari

	var luas = ((phi * r) * r)

	fmt.Println(luas)
}
