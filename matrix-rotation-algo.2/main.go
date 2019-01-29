package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var m, n uint64
	m, n = 4, 4

	data := make([][]uint64, m)

	var numero uint64
	for i := uint64(0); i < m; i++ {

		for y := uint64(0); y < n; y++ {

			data[i] = append(data[i], numero)
			numero++

		}
	}
	sizeM := len(data)
	layers := sizeM / 2
	fmt.Printf("%v \n", data)
	fmt.Printf("size mat: %v number layers: %v \n", sizeM, layers)

	for i := 0; i < layers; i++ {
		//first and last: identify the index position of the first and last rows and columns
		first := i
		last := sizeM - first - 1

		for y := first; y < last; y++ {
			offset := y - first

			fmt.Printf("Layer %d: first: %d, last: %d \n", layers, first, last)
			top := data[first][y]
			rightSide := data[y][last]
			bottom := data[last][last-offset]
			leftSide := data[last-offset][first]

			data[first][y] = leftSide
			data[y][last] = top
			data[last][last-offset] = rightSide
			data[last-offset][first] = bottom
		}

	}
	fmt.Println(data)

}
func after() {

	var m, n, r uint64
	fmt.Scanf("%v %v %v\n", &m, &n, &r)

	data := make([][]uint64, m)

	reader := bufio.NewReader(os.Stdin)
	for i := uint64(0); i < m; i++ {
		in, _ := reader.ReadString('\n')
		strArr := strings.Fields(in)
		for y := uint64(0); y < n; y++ {
			numero, _ := strconv.ParseUint(strArr[y], 10, 64)

			data[i] = append(data[i], numero)

		}
	}
	fmt.Printf("%v \n", data)
	fmt.Printf("%v \n%v", data[0][1], data[2][1])
}
