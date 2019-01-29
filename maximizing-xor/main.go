<<<<<<< HEAD
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
=======
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
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
