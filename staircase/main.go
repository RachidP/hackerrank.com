package main

import (
	"fmt"
)

func main() {

	var size int
	val := "#"
	fmt.Scanf("%v\n", &size)
	j := 0
	for i := 1; i <= size; i++ {

		j = size - i
		for j > 0 {

			fmt.Printf(" ")
			j--

		}
		fmt.Printf("%v\n", val)
		val += "#"

	}
}
