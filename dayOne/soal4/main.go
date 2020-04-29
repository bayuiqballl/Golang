package main

import (
	"fmt"
	"time"
)

func main() {
	sekarang := time.Now()

	fmt.Println("Hari ini : ", sekarang.Format("2020, April  28"))

	panjang := time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC)

	perbedaan := sekarang.Sub(panjang)

	hari := int(perbedaan.Hours() / 24)

	fmt.Printf("%d\n", hari)
}
