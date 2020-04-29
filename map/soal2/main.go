package main

import "fmt"

func search(value string) bool {

	takjil := [5]string{"kurma", "kolak", "es buah", "nasi", "roti"}

	for _, item := range takjil {
		if item == value {
			return true

		}
	}
	return false
}

func main() {
	key := search("buah")
	if key {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
