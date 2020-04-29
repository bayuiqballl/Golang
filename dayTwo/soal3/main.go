/*

Buatlah sebuah variabel untuk menyimpan informasi berupa array
Buatlah sistem yang mengenerate sebuah bilangan prima yang ditentukan dari x (bilangan min) dan y (bilangan max)
Hasil generate ditampung dalam sebuah variabel pada point (1)
Cetaklah ke dalam layar isi dari variabel pada point (1)

*/
package main

import "fmt"


func main(){
	slice := make([]int,0,20)
	y := 20

	for x := 1; x < y; x++ {
		cek := 0
		for i:= 1; i < y; i++ {
			if x % i == 0 {
				cek++
			}
		}
		if cek == 2 && y > 1 {
			slice = append(slice, x)
		}
	}
	fmt.Println(slice)
}




