<<<<<<< HEAD
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	var n uint64
	fmt.Scanf("%v\n", &t)

	for i := 1; i <= t; i++ {
		fmt.Scanf("%v\n", &n)

		counter := n

		isLouiseWinning := false

		if n == 1 {
			isLouiseWinning = false
			//fmt.Println("Richard")
		}

		for n > 1 {

			isLouiseWinning = !isLouiseWinning
			if !isPowerOfTwo(n) {
				counter = counter - getLargestPoweroftwo(float64(n))

			} else {
				tmp := n / 2
				counter -= tmp

			}
			n = counter

		}
		if isLouiseWinning {
			fmt.Println("Louise")
		} else {
			fmt.Println("Richard")

		}

	}

}

/* Function to check if x is power of 2*/
func isPowerOfTwo(x uint64) bool {
	//fmt.Println("isPower")
	return (x != 0) && ((x & (x - 1)) == 0)
}

func getLargestPoweroftwo(n float64) uint64 {

	var maxP float64 = -1

	var i float64

	for {
		p := math.Pow(2, i)
		i++

		if p > maxP {
			if p < n {
				maxP = p

			} else {
				break
			}

		}

	}
	return uint64(maxP)
}
=======
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	var n uint64
	fmt.Scanf("%v\n", &t)

	for i := 1; i <= t; i++ {
		fmt.Scanf("%v\n", &n)

		counter := n

		isLouiseWinning := false

		if n == 1 {
			isLouiseWinning = false
			//fmt.Println("Richard")
		}

		for n > 1 {

			isLouiseWinning = !isLouiseWinning
			if !isPowerOfTwo(n) {
				counter = counter - getLargestPoweroftwo(float64(n))

			} else {
				tmp := n / 2
				counter -= tmp

			}
			n = counter

		}
		if isLouiseWinning {
			fmt.Println("Louise")
		} else {
			fmt.Println("Richard")

		}

	}

}

/* Function to check if x is power of 2*/
func isPowerOfTwo(x uint64) bool {
	//fmt.Println("isPower")
	return (x != 0) && ((x & (x - 1)) == 0)
}

func getLargestPoweroftwo(n float64) uint64 {

	var maxP float64 = -1

	var i float64

	for {
		p := math.Pow(2, i)
		i++

		if p > maxP {
			if p < n {
				maxP = p

			} else {
				break
			}

		}

	}
	return uint64(maxP)
}
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
