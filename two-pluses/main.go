package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	up = iota
	left
	down
	right
)

var (
	nRows, nCln int
	mat         [][]int32
	m           = make(map[int][4]int)

	pluses [][]int
)

// Complete the twoPluses function below.
func twoPluses(grid []string) int32 {

	fmt.Println(grid)

	mat, nRows, nCln = transtoMat(grid)

	for i := 1; i < nRows-1; i++ {

		for j := 1; j < nCln-1; j++ {
			if mat[i][j] == 0 {
				continue
			}

			checkIsPlus(i, j)
		}
	}

	fmt.Println(pluses)
	return checkOverlay()

}

func checkOverlay() int32 {
	maxArea := math.MinInt32
	nPluses := len(pluses)
	var overlab bool
	for i := 0; i < nPluses; i++ {
		vv1 := pluses[i]
		for j := i + 1; j < len(pluses); j++ {
			overlab = false
			vv2 := pluses[j]

		Restart:
			for _, v1 := range vv1 {
				for _, v2 := range vv2 {
					if v1 == v2 {
						overlab = true
						break Restart
					}
				}

			}

			if !overlab {
				tmp := len(vv1) * len(vv2)
				if maxArea < tmp {

					maxArea = tmp
				}
			}

		}
	}
	fmt.Println("Max Area: ", maxArea)
	return int32(maxArea)

}

func checkIsPlus(row, cl int) {

	//neighbour := [4]int{up,left,down,right}
	neighbour := [4]int{0, 0, 0, 0}

	//concatenate numbers
	cncUp := concatenate(row-1, cl)
	cncLeft := concatenate(row, cl-1)
	//fmt.Println(row, cl, cnc)

	//check if the upper cell exist in the map, then update data
	if v, ok := m[cncUp]; ok {
		neighbour[up] = v[up] + 1
		neighbour[down] = v[down] - 1

	} else {

		//up: check if the up cell is 1
		if mat[row-1][cl] == 1 {

			neighbour[up] = 1

		}
		//count number of cells of 1 bellow the current cell
		for r := row + 1; r < nRows && (mat[r][cl] == 1); r++ {
			neighbour[down]++
		}

	}

	//check if the left cell exist in the map, then update data
	if v, ok := m[cncLeft]; ok {
		neighbour[left] = v[left] + 1
		neighbour[right] = v[right] - 1

	} else { //calcolate number elements on the right and left of the current cell

		//left: check if the LEFT cell is 1
		if mat[row][cl-1] == 1 {

			neighbour[left] = 1

		}
		//count number of cells on "1" on the RIGHT of the current cell
		for c := cl + 1; c < nCln && (mat[row][c] == 1); c++ {
			neighbour[right]++
		}

	}

	//add the
	center := concatenate(row, cl)
	m[center] = neighbour

	// area := getArea(neighbour)
	getPluses(neighbour, center)

	// fmt.Println(neighbour, area)

}

func getArea(mat [4]int) (int, int) {

	minAbs := math.MaxInt32 - 1

	for _, v := range mat {
		//check if is a plus (a plus can't contain a zero value)
		if v <= 0 {
			return 1, 0 //1 is the area of the current element
		}
		if v <= minAbs {
			minAbs = v

		}

	}

	return (minAbs * 4) + 1, minAbs
}
func getPluses(mat [4]int, center int) {

	minAbs := math.MaxInt32 - 1

	for _, v := range mat {
		//check if is a plus (a plus can't contain a zero value)
		// if v <= 0 {
		// 	return
		// }
		if v <= minAbs {
			minAbs = v

		}

	}

	plus := make([]int, 1+(minAbs*4))
	//PS: the order is important because sort the cells
	plus[len(plus)/2] = center
	for i := 0; i < minAbs; i++ {
		movTen := 10 * (i + 1)

		plus[i] = center - movTen //up

		plus[minAbs+i] = center - minAbs + i //left

		plus[(minAbs*2+1)+i] = center + (i + 1) //Right

		plus[(minAbs*3+1)+i] = center + movTen //Down
	}
	// for i := 0; i < minAbs; i++ {
	// 	movTen := 10 * (i + 1)

	// 	plus[i] = center - movTen //up

	// 	plus[minAbs+i] = center - minAbs + i //left

	// 	plus[(minAbs*2)+i] = center + (i + 1) //Right

	// 	plus[(minAbs*3)+i] = center + movTen //Down
	// }
	//fmt.Println(plus)
	pluses = append(pluses, plus)
	// //UP
	// for i := 0; i < minAbs; i++ {
	// 	plus[i] = center - 10*(i+1)

	// }
	// //LEFT
	// for i := 0; i < minAbs; i++ {
	// 	plus[i] = center - minAbs + i

	// }
	// //RIGHT
	// for i := 0; i < minAbs; i++ {
	// 	plus[i] = center + (i + 1)

	// }

	// //DOWN
	// for i := 0; i < minAbs; i++ {
	// 	plus[i] = center + 10*(i+1)

	// }
	fmt.Println("plus:", plus)

	//return (minAbs * 4) + 1, minAbs
}

//concatenate two numbers ex a=10, b=20 it will return c= 1020
func concatenate(a, b int) int {

	result, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	checkError(err)
	return result
}

func transtoMat(grid []string) (res [][]int32, nRows, nCln int) {
	nRows = len(grid)

	for _, v := range grid {
		nCln = 0
		row := make([]int32, 0)
		//fmt.Println(v)
		for _, s := range v {
			nCln++
			if s == 'G' {
				row = append(row, 1)
				continue
			}
			row = append(row, 0)

		}
		res = append(res, row)
		//fmt.Printf("row = %v \n", row)

	}
	// fmt.Println(nRows, nCln)
	// fmt.Println(res)
	return
}

func main() {

	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("OUTPUT_PATH")
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)
	_ = m
	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := twoPluses(grid)

	fmt.Fprintf(writer, "%d\n", result)
	//fmt.Println(pluses)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func example() {

	guestList := []string{"bill", "jill", "joan"}
	arrived := []string{"sally", "jill", "joan"}

CheckList:
	for _, guest := range guestList {
		for _, person := range arrived {
			fmt.Printf("Guest[%s] Person[%s]\n", guest, person)

			if person == guest {
				fmt.Printf("Let %s In\n", person)
				break CheckList
			}
		}
	}
}
