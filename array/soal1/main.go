package main

import "fmt"

func main() {
	array1 := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	array2 := [20]int{1, 2, 30, 4, 5, 6, 7, 8, 9, 110, 111, 112, 113, 14, 15, 16, 17, 18, 19, 20}
	diff := []int{}

	for i := 0; i < len(array1); i++ {
		status := false
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				status = true

			}

		}
		if !status {
			diff = append(diff, array1[i])
		}

	}

	for i := 0; i < len(array2); i++ {
		status := false
		for j := 0; j < len(array1); j++ {
			if array2[i] == array1[j] {
				status = true

			}

		}
		if !status {

			diff = append(diff, array2[i])
		}
	}

	fmt.Println(diff)
}
