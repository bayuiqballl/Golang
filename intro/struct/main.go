package main

import "fmt"

type Laptop struct {
	merk string
	hdd  int
	ram int	
}

func (laptop *Laptop) changeMerk(newMerk string ){
	laptop.merk = newMerk
}

func main() {
	l := Laptop{"ASUS", 512, 2048}
	// l  = Laptop{"Acer"}
	l.changeMerk("ACER")

	fmt.Println(l)

}
