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
