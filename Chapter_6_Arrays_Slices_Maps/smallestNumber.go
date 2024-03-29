package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minimum := x[0]
	for _, element := range x {
		if element < minimum {
			minimum = element
		}
	}

	fmt.Println(minimum)
}
