package main

import "fmt"

type Produk struct {
	nama  string
	harga int
	stok  int
}

func main() {
	minimarket := make(map[int]Produk)
	minimarket[1] = Produk{nama: "Roti Tawar", harga: 13000, stok: 15}
	minimarket[2] = Produk{nama: "Sabun", harga: 1300, stok: 14}
	minimarket[3] = Produk{nama: "Aqua", harga: 1000, stok: 10}
	minimarket[4] = Produk{nama: "Kurma", harga: 3000, stok: 10}

	for _, item := range minimarket {
		if item.stok <= 10 {
			fmt.Println("Nama Barang: ", item.nama, "\nHarga Barang: Rp.", item.harga, "\nStok Barang: ", item.stok)
		}
	}

}
