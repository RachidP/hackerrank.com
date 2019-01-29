<<<<<<< HEAD
package main

import (
	"fmt"
)

func main() {
	var numQ int
	var l, r, tot uint64

	fmt.Scanf("%v\n", &numQ)
	for i := 1; i <= numQ; i++ {

		fmt.Scanf("%v %v\n", &l, &r)
		tot = calcola(r) ^ calcola(l-1)

		fmt.Println(tot)

	}

}
func calcola(x uint64) uint64 {
	a := uint64(x % 8)
	if a == 0 || a == 1 {
		return x
	}
	if a == 2 || a == 3 {
		return 2
	}
	if a == 4 || a == 5 {
		return x + 2
	}
	if a == 6 || a == 7 {
		return 0
	}
	return 0

}
=======
package main

import (
	"fmt"
)

func main() {
	var numQ int
	var l, r, tot uint64

	fmt.Scanf("%v\n", &numQ)
	for i := 1; i <= numQ; i++ {

		fmt.Scanf("%v %v\n", &l, &r)
		tot = calcola(r) ^ calcola(l-1)

		fmt.Println(tot)

	}

}
func calcola(x uint64) uint64 {
	a := uint64(x % 8)
	if a == 0 || a == 1 {
		return x
	}
	if a == 2 || a == 3 {
		return 2
	}
	if a == 4 || a == 5 {
		return x + 2
	}
	if a == 6 || a == 7 {
		return 0
	}
	return 0

}
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
