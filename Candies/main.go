package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var m []int64

// Complete the candies function below.
func candies(n int32, arr []int32) int64 {
	//	fmt.Println(arr)
	m = make([]int64, n)
	var tot int64
	sizeA := int(n)

	m[0] = 1
	for i := 1; i < sizeA; i++ {
		if arr[i] > arr[i-1] {
			m[i] = m[i-1] + 1
		} else {
			m[i] = 1
		}
	}

	tot = m[sizeA-1]

	for i := sizeA - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] && m[i] <= m[i+1] {
			m[i] = m[i+1] + 1

		}
		tot += m[i]

	}
	//	fmt.Println(m)
	fmt.Println(tot)
	return tot

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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := candies(n, arr)

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
