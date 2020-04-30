package main

import "fmt"

func main() {
	nama := make(map[int]string)
	nama[0] = ("Roti Tawar")
	nama[1] = ("Sabun")
	nama[2] = ("Aqua")
	nama[3] = ("Kurma")

	harga := make(map[int]int)
	harga[0] = (13000)
	harga[1] = (3000)
	harga[2] = (5000)
	harga[3] = (50000)

	stok := make(map[int]int)
	stok[0] = (5)
	stok[1] = (4)
	stok[2] = (10)
	stok[3] = (10)

	for i, item := range stok {
		if item < 10 {
			fmt.Println("Nama Barang : ", nama[i])
			fmt.Println("Harga  :", harga[i])
			fmt.Println("Stok  :", stok[i])
		}
	}

}
