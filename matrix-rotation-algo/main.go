package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var m, n, r uint64
	fmt.Scanf("%v %v %v\n", &m, &n, &r)
	sizeVec := m * n
	data := make([]uint64, 0, sizeVec)

	reader := bufio.NewReader(os.Stdin)
	for i := uint64(0); i < m; i++ {
		in, _ := reader.ReadString('\n')
		strArr := strings.Fields(in)
		for y := uint64(0); y < n; y++ {
			numero, _ := strconv.ParseUint(strArr[y], 10, 64)
			data = append(data, numero)

		}
	}

	fmt.Println(data)
	printMat(data, m)
	layers := getMax(m, n) / 2

	if sizeVec >= 1 {
		for index := uint64(0); index < r; index++ {

			tmp := data[0]
			crescente := true
			//data[n] = data[0]

			var layer, first, i uint64
			last := sizeVec - first - 1
			lenRow := m
			offset := lenRow

			lowerBound := first + lenRow
			upperBound := last - lenRow

			for {
				// calcolate offset

				if crescente {

					i += offset
					data[i], tmp = tmp, data[i]

					if i >= last {
						crescente = !crescente

						offset = lenRow
						continue

					}
				} else {

					i -= offset
					data[i], tmp = tmp, data[i]

					if i == first {

						layer++
						if layer < layers {
							first += (m + 1)
							last = sizeVec - first - 1
							i = first
							tmp = data[first]
							crescente = !crescente
							offset = lenRow
							lowerBound = first + lenRow
							upperBound = last - lenRow
							continue

						} else {

							break
						}

					}
				}

				if (i < lowerBound) || (i > upperBound) {
					offset = 1
				} else {

					offset = lenRow
				}

			}
		}
	}
	fmt.Println(data)
	printMat(data, m)

}
func getMax(x1, x2 uint64) uint64 {
	if x1 >= x2 {
		return x1
	}
	return x2

}

func printMat(a []uint64, r uint64) {
	i := uint64(1)
	for _, v := range a {
		if r < i {
			fmt.Printf("\n")
			i = 1
		}
		fmt.Printf("%v ", v)
		i++

	}
	fmt.Printf("\n")
}
