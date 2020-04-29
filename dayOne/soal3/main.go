package main

import "fmt"

/* program fizzbuzz
jika habis dibagi 3 Fizz,
jika habis dibagi 5 Buzz
jika habis dibagi 15 FizzBuzz


*/
func main() {

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}

}
