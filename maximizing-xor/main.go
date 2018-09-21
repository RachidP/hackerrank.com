package main

import (
	"fmt"
)

func main() {

	var l, r int
	max := -1
	fmt.Scanf("%v\n%v\n", &l, &r)

	for i := l; i <= r; i++ {

		for j := i; j <= r; j++ {
			tmp := (i ^ j)
			if max < tmp {
				max = tmp
			}

		}
	}
	fmt.Println(max)
}
