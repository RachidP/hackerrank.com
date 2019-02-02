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
	m           = make(map[string][4]int)

	pluses [][]string
)

// Complete the twoPluses function below.
func twoPluses(grid []string) int32 {

	fmt.Println(grid)

	mat, nRows, nCln = transtoMat(grid)
	//fmt.Println(mat)
	for i := 1; i < nRows-1; i++ {

		for j := 1; j < nCln-1; j++ {
			if mat[i][j] == 0 {
				continue
			}

			checkIsPlus(i, j)
		}
	}

	//fmt.Println(pluses)
	return checkOverlay()

}

func checkOverlay() int32 {
	maxArea := math.MinInt32
	nPluses := len(pluses)
	var overlab bool
	for i := 0; i < nPluses; i++ {
		vv1 := pluses[i]
		for j := i + 1; j < nPluses; j++ {
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
					//fmt.Println("newMax: ", vv1, vv2, " area ", tmp)
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

	//check if the LEFT/RIGHT cell exist in the map, then update data
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
	getPluses(neighbour, row, cl)

}

func getPluses(mat [4]int, row, cl int) {

	minAbs := math.MaxInt32 - 1

	for _, v := range mat {

		if v <= minAbs {
			minAbs = v

		}

	}
	for minAbs >= 0 {

		plus := make([]string, 1+(minAbs*4))

		plus[len(plus)/2] = concatenate(row, cl)
		for i := 0; i < minAbs; i++ {
			movTen := (i + 1)

			plus[i] = concatenate(row-minAbs+i, cl) //up

			plus[minAbs+i] = concatenate(row, cl-minAbs+i) //left

			plus[(minAbs*2+1)+i] = concatenate(row, cl+movTen) //Right

			plus[(minAbs*3+1)+i] = concatenate(row+movTen, cl) //Down

		}

		pluses = append(pluses, plus)
		minAbs--

	}
	// fmt.Println("plus:", plus)

}

//concatenate two numbers ex a=10, b=20 it will return c= 1020
func concatenate(row, cl int) string {

	return fmt.Sprintf("%d.%d", row, cl)
}

func transtoMat(grid []string) (res [][]int32, nRows, nCln int) {
	nRows = len(grid)

	for _, v := range grid {
		nCln = 0
		row := make([]int32, 0)

		for _, s := range v {
			nCln++
			if s == 'G' {
				row = append(row, 1)
				continue
			}
			row = append(row, 0)

		}
		res = append(res, row)

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
