<<<<<<< HEAD
package main

import (
	"fmt"
)

func main() {

	const (
		c0 = iota // c0 == 0 print Actual State of the matrix
		c1 = iota // c1 == 1 print all elements of the matrix equal to "O"
		c2 = iota // c2 == 2 print the matrix after trasformation to the right output O,E->O, c->.
		c3 = iota // c3 ==3 rebuild the matrix from O,E->O, C->.
	)
	var m, n, t int
	//fmt.Scanf("%v %v %v\n", &m, &n, &t)
	m, n, t = 6, 7, 3
	sizeVec := m * n
	data := make([]string, 0, sizeVec)
	inputT := []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}
	//1. Initially, Bomberman arbitrarily plants bombs in some of the cells.
	//var numero uint64
	//	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		tmp := inputT[i]
		//in, _ := reader.ReadString('\n')
		for _, r := range tmp {
			c := string(r)
			data = append(data, c)

		}

	}
	//printMat(data, n, c0)

	switch {
	case t == 1:
		printMat(data, n, c0)
	case t%2 == 0:
		printMat(data, n, c1)
	case t%4 == 3: //the grid will be the some one in 3rd second
		change(data, n)
		//	printMat(data, n, c0)
		printMat(data, n, c2)
	case t%4 == 1: //the grid will be the some one in 5th second
		change(data, n)
		printMat(data, n, c3)
		change(data, n)
		printMat(data, n, c2)
	}

}

func change(a []string, n int) {
	lenA := len(a)
	var tmp int
	/*
		O: current Bomb -> B
		.: close to bomb -> E
		.: Change status -> C
	*/
	for i := 0; i < lenA; i++ {
		if a[i] == "O" {

			//right cell
			tmp = i + 1
			if tmp < lenA && a[tmp] != "O" {
				a[tmp] = "E"

			}
			//left cell
			tmp = i - 1
			if tmp >= 0 && i%n != 0 && a[tmp] != "O" {
				a[tmp] = "E"
			}
			// down cell
			tmp = i + n
			if tmp < lenA && a[tmp] != "O" {
				a[tmp] = "E"
			}

			// up cell
			tmp = i - n
			if tmp >= 0 && a[tmp] != "O" {
				a[tmp] = "E"
			}

		} else if a[i] == "." {

			a[i] = "C"

		}

	}

}

func printMat(a []string, r int, t int) {
	dim := len(a)
	for i := 0; i < dim; i++ {

		v := a[i]
		switch t {
		case 0:
			fmt.Printf("%v", v)
		case 1:
			fmt.Printf("O")
		case 2:
			if v == "E" || v == "O" {
				fmt.Printf(".")
			} else {
				fmt.Printf("O")
			}
		case 3:
			if v == "E" || v == "O" {
				a[i] = "."
			} else {
				a[i] = "O"
			}

		}
		if (i+1)%(r) == 0 && t != 3 {
			fmt.Printf("\n")
		}

	}
	//fmt.Printf("\n**************************\n")

}
=======
package main

import (
	"fmt"
)

func main() {

	const (
		c0 = iota // c0 == 0 print Actual State of the matrix
		c1 = iota // c1 == 1 print all elements of the matrix equal to "O"
		c2 = iota // c2 == 2 print the matrix after trasformation to the right output O,E->O, c->.
		c3 = iota // c3 ==3 rebuild the matrix from O,E->O, C->.
	)
	var m, n, t int
	//fmt.Scanf("%v %v %v\n", &m, &n, &t)
	m, n, t = 6, 7, 3
	sizeVec := m * n
	data := make([]string, 0, sizeVec)
	inputT := []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}
	//1. Initially, Bomberman arbitrarily plants bombs in some of the cells.
	//var numero uint64
	//	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		tmp := inputT[i]
		//in, _ := reader.ReadString('\n')
		for _, r := range tmp {
			c := string(r)
			data = append(data, c)

		}

	}
	//printMat(data, n, c0)

	switch {
	case t == 1:
		printMat(data, n, c0)
	case t%2 == 0:
		printMat(data, n, c1)
	case t%4 == 3: //the grid will be the some one in 3rd second
		change(data, n)
		//	printMat(data, n, c0)
		printMat(data, n, c2)
	case t%4 == 1: //the grid will be the some one in 5th second
		change(data, n)
		printMat(data, n, c3)
		change(data, n)
		printMat(data, n, c2)
	}

}

func change(a []string, n int) {
	lenA := len(a)
	var tmp int
	/*
		O: current Bomb -> B
		.: close to bomb -> E
		.: Change status -> C
	*/
	for i := 0; i < lenA; i++ {
		if a[i] == "O" {

			//right cell
			tmp = i + 1
			if tmp < lenA && a[tmp] != "O" {
				a[tmp] = "E"

			}
			//left cell
			tmp = i - 1
			if tmp >= 0 && i%n != 0 && a[tmp] != "O" {
				a[tmp] = "E"
			}
			// down cell
			tmp = i + n
			if tmp < lenA && a[tmp] != "O" {
				a[tmp] = "E"
			}

			// up cell
			tmp = i - n
			if tmp >= 0 && a[tmp] != "O" {
				a[tmp] = "E"
			}

		} else if a[i] == "." {

			a[i] = "C"

		}

	}

}

func printMat(a []string, r int, t int) {
	dim := len(a)
	for i := 0; i < dim; i++ {

		v := a[i]
		switch t {
		case 0:
			fmt.Printf("%v", v)
		case 1:
			fmt.Printf("O")
		case 2:
			if v == "E" || v == "O" {
				fmt.Printf(".")
			} else {
				fmt.Printf("O")
			}
		case 3:
			if v == "E" || v == "O" {
				a[i] = "."
			} else {
				a[i] = "O"
			}

		}
		if (i+1)%(r) == 0 && t != 3 {
			fmt.Printf("\n")
		}

	}
	//fmt.Printf("\n**************************\n")

}
>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
