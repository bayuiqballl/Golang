package main

import "fmt"

type Sepeda struct {
	ban, gigi int
	warna     string
}

func (s *Sepeda) time(menit float32) float32 {

	waktu := menit * 2.5

	return waktu
}

func main() {
	item := make(map[int]Sepeda)

	item[0] = Sepeda{ban: 2, gigi: 2, warna: "Hijau"}
	item[1] = Sepeda{ban: 2, gigi: 5, warna: "Merah"}
	item[2] = Sepeda{ban: 2, gigi: 1, warna: "Hitam"}
	item[3] = Sepeda{ban: 2, gigi: 4, warna: "Biru"}
	item[4] = Sepeda{ban: 2, gigi: 3, warna: "Putih"}

	fmt.Print("List Sepeda\n")

	for index, i := range item {
		fmt.Println("============Deskripsi Sepeda ke :", index+1, " ===============")
		fmt.Println("Warna  = ", i.warna)
		fmt.Println("Jumlah Ban  = ", i.ban)
		fmt.Println("Jumlah Gear  = ", i.gigi)
		fmt.Println("Waktu tempuh nya = ", i.time(20), "menit")
	}
}
