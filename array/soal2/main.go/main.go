package main

import "fmt"

func main() {
	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var genap, ganjil, prima []int

	fmt.Println("array ori adalah : ", array)

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			genap = append(genap, array[i])
		}
	}
	fmt.Println("array genap adalah : ", genap, "jumlah: ", len(genap))

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			ganjil = append(ganjil, array[i])
		}
	}
	fmt.Println("array ganjil adalah : ", ganjil, "jumlah: ", len(ganjil))

	for i := 1; i < len(array); i++ {
		count := 0

		for y := 1; y <= array[i]; y++ {
			if array[i]%y == 0 {
				count++
			}
		}
		if count == 2 && array[i] > 1 {
			prima = append(prima, array[i])

		}
	}

	fmt.Println("array prima adalah : ", prima, "jumlah: ", len(prima))

}
