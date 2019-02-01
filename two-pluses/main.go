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
	maxArea     = -100
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
	checkOverlay()
	return -1
}

func checkOverlay() {
	for k, v := range m {
		area, length := getArea(v)
		if area == 1 {
			continue
		}
		fmt.Println(k, v, area, length)
		buildPlus(k, length)
	}

}

func buildPlus(k, lenght int) {

}

func checkIsPlus(row, cl int) {
	// //check if there is a
	// count := 4
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

		//left: check if the left cell is 1
		if mat[row][cl-1] == 1 {

			neighbour[left] = 1
		}
		//count number of cells on "1" on the right of the current cell
		for c := cl + 1; c < nCln && (mat[row][c] == 1); c++ {
			neighbour[right]++
		}
	}

	//add the
	m[concatenate(row, cl)] = neighbour

	// area := getArea(neighbour)

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
