package main

import "fmt"



func main(){

	a := [2]int{1, 2}
	b := [...]int{1,2}
	c := [2]int{1,3}

	fmt.Println(a == c)
	fmt.Println(a == b)
	fmt.Println(b == c)
}