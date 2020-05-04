package main

import "fmt"

type Maju interface {
	cepat() float64
	lambat() float64
}

type sepeda struct {
	ban, gigi int
	warna     string
	menit     float64
}

type motor struct {
	ban, gigi int
	warna     string
	menit     float64
}

func (s sepeda) cepat() float64 {

	waktu := s.menit * 2.5 * 2

	return waktu
}

func (s sepeda) lambat() float64 {

	waktu := s.menit * 2.5 * 0.5

	return waktu
}

func (m motor) cepat() float64 {

	waktu := m.menit * 5 * 2

	return waktu
}

func (m motor) lambat() float64 {

	waktu := m.menit * 5 * 0.5

	return waktu
}

func hitung(b Maju) {
	fmt.Println("cepat :", b.cepat())
	fmt.Println("lambat :", b.lambat())
}

func main() {
	kecepatanSepeda := sepeda{2, 2, "merah", 20}
	fmt.Println(kecepatanSepeda)
	fmt.Println("waktu tempuh : ")
	hitung(kecepatanSepeda)
}
