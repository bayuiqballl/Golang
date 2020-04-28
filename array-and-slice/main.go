package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	var y [8]string
	fmt.Println(y)

	var names [4]string
	names[0] = "Lutfi"
	names[1] = "Arya"
	names[2] = "Bella"
	names[3] = "rahmat"
	fmt.Println(names, names[1])

	// menampilkan array
	a := [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//slice
	r := [...]int{1, 2, 3, 4}
	fmt.Println(r)
	fmt.Println(r[0:2])
	fmt.Println(r[1:2])
	fmt.Println(r[2:])

}
