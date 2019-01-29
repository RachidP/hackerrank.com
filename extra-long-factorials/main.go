<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var numQ int64

	fmt.Scanf("%v\n", &numQ)
	//fmt.Println(calcola(numQ))
	//	x := big.NewInt(numQ)
	fmt.Println(calcola(numQ))
}

func calcola(x int64) *big.Int {
	tmp := big.NewInt(x)
	if x <= 0 {
		return big.NewInt(0)

	}
	if x == 1 {

		return big.NewInt(1)
	}
	return tmp.Mul(tmp, calcola(x-1))

}
=======
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var numQ int64

	fmt.Scanf("%v\n", &numQ)
	//fmt.Println(calcola(numQ))
	//	x := big.NewInt(numQ)
	fmt.Println(calcola(numQ))
}

func calcola(x int64) *big.Int {
	tmp := big.NewInt(x)
	if x <= 0 {
		return big.NewInt(0)

	}
	if x == 1 {

		return big.NewInt(1)
	}
	return tmp.Mul(tmp, calcola(x-1))

}
<<<<<<< HEAD
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
=======
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
